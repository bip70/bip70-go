package bip70

const (
	// PkiTypeNone indicates that the request has no
	// pki data or signature
	PkiTypeNone = "none"

	// PkiTypeX509Sha256 indicates the hashing algorithm
	// to be taken with the public key type to produce
	// the signature algorithm.
	PkiTypeX509Sha256 = "x509+sha256"

	// PkiTypeX509Sha1 indicates the hashing algorithm
	// to be taken with the public key type to produce
	// the signature algorithm.
	PkiTypeX509Sha1 = "x509+sha1"

	// MimeTypePaymentRequest is the MIME type, set as
	// Accept: header to download a request, and set
	// by the server as Content-Type in a successful
	// response
	MimeTypePaymentRequest = "application/bitcoin-paymentrequest"

	// MimeTypePayment is the MIME type, set as
	// Content-Type: header when submitting a payment
	// to the server
	MimeTypePayment = "application/bitcoin-payment"

	// MimeTypePaymentAck is the MIME type, set as
	// by the server as Content-Type in a successful
	// response
	MimeTypePaymentAck = "application/bitcoin-paymentack"
)
