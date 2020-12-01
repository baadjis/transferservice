package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	
	"github.com/baadjis/transferservice/graph/generated"
	"github.com/baadjis/transferservice/graph/model"
)

func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.TransactionInput) (*model.Transaction, error) {
	sender,_:=r.AddCustomer(ctx,*input.Sender)
	receiver,_:=r.AddCustomer(ctx,*input.Receiver)
	details,_:=r.AddDetails(ctx,*input.TransactionDetails)
    transaction:= model.Transaction{
	   Txcode :input.Txcode,
	   SenderAgentID:input.SenderAgentID,
	   ReceiverAgentID:input.ReceiverAgentID,
	   Sender:sender.ID,
	   Receiver:receiver.ID,
	   TransactionDetails:details.ID,
	   Status:input.Status}
   fmt.Println(r.DB)
   err := r.DB.Create(&transaction).Error
   if err != nil {
	   return nil, err
   }
   return &transaction,nil
}

func (r *mutationResolver) UpdateTransaction(ctx context.Context, txcode string, input model.TransactionInput) (*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTransaction(ctx context.Context, txcode string) (bool, error) {
	r.DB.Where("Txcode = ?", txcode).Delete(&model.TransactionDetails{})
    r.DB.Where("Txcode= ?", txcode).Delete(&model.Transaction{})
	return true, nil
}

func (r *mutationResolver) ConfirmReception(ctx context.Context, txcode string, input model.ReceptionInput) (bool, error) {
	
	r.DB.Model(model.Transaction{}).Where("Txcode= ?", txcode).Updates(model.Transaction{ReceiverAgentID:input.ReceiverAgentID,Status:input.Status})
	return true ,nil
}

func (r *mutationResolver) AddCustomer(ctx context.Context, input model.CustomerInput) (*model.Customer, error) {
	customer:=model.Customer{
		Firstname :input.Firstname,
		Lastname  :input.Lastname,
		Phone     :input.Phone ,
		Email     :input.Email,
		Country  :input.Country}

        r.DB.Where(model.Customer{Phone:input.Phone}).FirstOrCreate(&customer)
  
		return &customer,nil
}

func (r *mutationResolver) ChangeReceiver(ctx context.Context, txcode string, input model.CustomerChanges) (*model.Customer, error) {
	var transaction model.Transaction
	r.DB.Where("Txcode= ?", txcode).First(&transaction)
	customerid:=transaction.Receiver
	var customer model.Customer
    r.DB.First(&customer,customerid)
	return &customer,nil
}

func (r *mutationResolver) DeletCustomer(ctx context.Context, id string) (bool, error) {
	r.DB.Where("id= ?", id).Delete(&model.Transaction{})
	return true, nil
}

func (r *mutationResolver) AddDetails(ctx context.Context, input model.TransactionDetailsInput) (*model.TransactionDetails, error) {
	dt:= model.TransactionDetails{
		Txcode         :input.Txcode,
		SentAmount      :input.SentAmount,
		SentCurrency     :input.SentCurrency,
		ReceivedAmount   :input.ReceivedAmount,
		ReceivedCurrency :input.ReceivedCurrency,
		Xchange          :input.Xchange,
		Fees             :input.Fees,
		PaymentMode      :input.PaymentMode,
		ReceptionMode    :input.ReceptionMode}


        r.DB.Where(model.TransactionDetails{Txcode:input.Txcode}).FirstOrCreate(&dt)
  
		return &dt,nil
}

func (r *mutationResolver) DeleteDetails(ctx context.Context, txcode string) (bool, error) {
	r.DB.Where("Txcode = ?", txcode).Delete(&model.TransactionDetails{})
    r.DB.Where("Txcode= ?", txcode).Delete(&model.Transaction{})
	return true, nil
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	r.DB.Find(&transactions)
    return transactions,nil
}

func (r *queryResolver) Transaction(ctx context.Context, txcode string) (*model.Transaction, error) {
	var transaction model.Transaction
	r.DB.Joins(" JOIN transaction_details ON transaction_details.txcode = transactions.txcode").
	Joins(" JOIN customers ON customers.id = transactions.sender").
	Where("transactions.txcode = ?" ,txcode).
	First(&transaction)

    return &transaction,nil
}

func (r *queryResolver) CustomerReceptions(ctx context.Context, id string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CustomerExpeditions(ctx context.Context, id string) ([]*model.Transaction, error) {
	panic(fmt.Errorf("not implemented"))
}


// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
