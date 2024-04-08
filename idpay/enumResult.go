package idpay

import "fmt"

// وضعیت تراکنش
// TransactionStatus
var (
	// TransactionStatusNotDone Payment has not been done
	// پرداخت انجام نشده است 1
	TransactionStatusNotDone = newTransactionStatus(1, "Payment has not been done", "پرداخت انجام نشده است")
	// TransactionStatusFailed Payment failed
	// پرداخت ناموفق بوده است 2
	TransactionStatusFailed = newTransactionStatus(2, "Payment failed", "پرداخت ناموفق بوده است")
	// TransactionStatusError Error occurred
	// خطا رخ داده است 3
	TransactionStatusError = newTransactionStatus(3, "Error occurred", "خطا رخ داده است")
	// TransactionStatusBlocked Blocked
	// بلوکه شده 4
	TransactionStatusBlocked = newTransactionStatus(4, "Blocked", "بلوکه شده")
	// TransactionStatusRefundedToPayer Refunded to the payer
	// برگشت به پرداخت کننده 5
	TransactionStatusRefundedToPayer = newTransactionStatus(5, "Refunded to the payer", "برگشت به پرداخت کننده")
	// TransactionStatusSystematicallyRefunded Systematically refunded
	// برگشت خورده سیستمی 6
	TransactionStatusSystematicallyRefunded = newTransactionStatus(6, "Systematically refunded", "برگشت خورده سیستمی")
	// TransactionStatusCanceled Payment canceled
	// انصراف از پرداخت 7
	TransactionStatusCanceled = newTransactionStatus(7, "Payment canceled", "انصراف از پرداخت")
	// TransactionStatusRedirectedToGateway Redirected to the payment gateway
	// به درگاه پرداخت منتقل شد 8
	TransactionStatusRedirectedToGateway = newTransactionStatus(8, "Redirected to the payment gateway", "به درگاه پرداخت منتقل شد")
	// TransactionStatusWaitingForConfirmation Waiting for payment confirmation
	// در انتظار تایید پرداخت 10
	TransactionStatusWaitingForConfirmation = newTransactionStatus(10, "Waiting for payment confirmation", "در انتظار تایید پرداخت")
	// TransactionStatusConfirmed Payment confirmed
	// پرداخت تایید شده است 100
	TransactionStatusConfirmed = newTransactionStatus(100, "Payment confirmed", "پرداخت تایید شده است")
	// TransactionStatusAlreadyConfirmed Payment has already been confirmed
	// پرداخت قبلا تایید شده است 101
	TransactionStatusAlreadyConfirmed = newTransactionStatus(101, "Payment has already been confirmed", "پرداخت قبلا تایید شده است")
	// TransactionStatusPaidToReceiver Paid to the receiver
	// به دریافت کننده واریز شد 200
	TransactionStatusPaidToReceiver = newTransactionStatus(200, "Paid to the receiver", "به دریافت کننده واریز شد")
	// unknownTransactionStatus Unknown transaction status
	// وضعیت تراکنش نامشخص است -1
	unknownTransactionStatus = newTransactionStatus(-1, "Unknown transaction status", "وضعیت تراکنش نامشخص است")
)

func GetTransactionStatus(status string) *TransactionStatus {
	switch status {
	case TransactionStatusNotDone.StatusStr:
		return TransactionStatusNotDone
	case TransactionStatusFailed.StatusStr:
		return TransactionStatusFailed
	case TransactionStatusError.StatusStr:
		return TransactionStatusError
	case TransactionStatusBlocked.StatusStr:
		return TransactionStatusBlocked
	case TransactionStatusRefundedToPayer.StatusStr:
		return TransactionStatusRefundedToPayer
	case TransactionStatusSystematicallyRefunded.StatusStr:
		return TransactionStatusSystematicallyRefunded
	case TransactionStatusCanceled.StatusStr:
		return TransactionStatusCanceled
	case TransactionStatusRedirectedToGateway.StatusStr:
		return TransactionStatusRedirectedToGateway
	case TransactionStatusWaitingForConfirmation.StatusStr:
		return TransactionStatusWaitingForConfirmation
	case TransactionStatusConfirmed.StatusStr:
		return TransactionStatusConfirmed
	case TransactionStatusAlreadyConfirmed.StatusStr:
		return TransactionStatusAlreadyConfirmed
	case TransactionStatusPaidToReceiver.StatusStr:
		return TransactionStatusPaidToReceiver
	default:
		return unknownTransactionStatus
	}
}

type TransactionStatus struct {
	Status    int    `json:"status"`
	StatusStr string `json:"status_str"`
	English   string `json:"en"`
	Persian   string `json:"fa"`
}

func newTransactionStatus(
	status int,
	en string,
	fa string,
) *TransactionStatus {
	return &TransactionStatus{
		Status:    status,
		StatusStr: fmt.Sprintf("%d", status),
		English:   en,
		Persian:   fa,
	}
}
