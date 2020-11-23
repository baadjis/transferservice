package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/baadjis/transferservice/graph/generated"
	"github.com/baadjis/transferservice/graph/model"
)



func mapCustomer(input model.CustomerInput)*model.Customer{
	custumer:=model.Customer{
		Firstname :input.Firstname,
		Lastname  :input.Lastname,
		Phone     :input.Phone ,
		Email     :input.Email,
		Country  :input.Country}
		return &custumer
}



func mapTxDetails(input model.TransactionDetailsInput) *model.TransactionDetails{
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
		return &dt
}
func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.TransactionInput) (*model.Transaction, error) {
	transaction:= model.Transaction{
		Txcode :input.Txcode,
		SenderAgentID:input.SenderAgentID,
		ReceiverAgentID:input.ReceiverAgentID,
		Sender:mapCustomer(*input.Sender),
		Receiver:mapCustomer(*input.Receiver),
		TransactionDetails:mapTxDetails(*input.TransactionDetails),
		Status:input.Status}
	fmt.Println(r.DB)
	err := r.DB.Create(&transaction).Error
    if err != nil {
        return nil, err
    }
	return &transaction,nil
}

func (r *mutationResolver) UpdateTransaction(ctx context.Context, txcode string, input *model.TransactionInput) (*model.Transaction, error) {
	transaction:= model.Transaction{
		Txcode :txcode,
		SenderAgentID:input.SenderAgentID,
		ReceiverAgentID:input.ReceiverAgentID,
		Sender:mapCustomer(*input.Sender),
		Receiver:mapCustomer(*input.Receiver),
		TransactionDetails:mapTxDetails(*input.TransactionDetails),
		Status:input.Status}
		err := r.DB.Save(&transaction).Error
    if err != nil {
        return nil, err
    }
	return &transaction,err
}

func (r *mutationResolver) DeleteTransaction(ctx context.Context, txcode string) (bool, error) {
	r.DB.Where("Txcode = ?", txcode).Delete(&model.TransactionDetails{})
    r.DB.Where("Txcode= ?", txcode).Delete(&model.Transaction{})
	return true, nil
}

func (r *queryResolver) Transactions(ctx context.Context) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	r.DB.Preload("Sender").Preload("Receiver").Preload("TransactionDetails").Find(&transactions)
    return transactions,nil
}

func (r *queryResolver) Transation(ctx context.Context, txcode string) (*model.Transaction, error) {
	var transaction *model.Transaction
	r.DB.Where("Txcode = ?",txcode).Preload("Sender").Preload("Receiver").Preload("TransactionDetails").Find(&transaction)

    return transaction,nil
	
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
