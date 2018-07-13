package transaction

type TransactionCallbackBase struct {
	transactionArray
}

func (this *TransactionCallbackBase) OnCommitBegin(transaction ITransaction) error {
	return nil
}

func (this *TransactionCallbackBase) OnCommitOver(transaction ITransaction) error {
	return nil
}

func (this *TransactionCallbackBase) OnSuccess(transaction ITransaction) {
}

func (this *TransactionCallbackBase) OnError(transaction ITransaction, err error) {
	if transaction.IsCommit() && err != nil {
		for _, ops := range transaction.GetData().([]ITransaction) {
			this.callOnError(ops, err)
		}
	}
}
