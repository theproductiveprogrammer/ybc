package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/mailgun/mailgun-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading configuration!")
		fmt.Println(err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getPaid(cfg, r, w)
	})
	fmt.Println("Server starting on :5463...")
	http.ListenAndServe(":5463", nil)
}

func loadConfig() (config, error) {
	var cfg config
	_, err := toml.DecodeFile("ybc.cfg", &cfg)
	if err != nil {
		return cfg, err
	}

	b, err := ioutil.ReadFile("mail.txt")
	if err != nil {
		return cfg, err
	}
	cfg.MailTxt = string(b)

	b, err = ioutil.ReadFile("mail.html")
	if err != nil {
		return cfg, err
	}
	cfg.MailHtml = string(b)

	return validateConfigx1(cfg)
}

func getPaid(cfg config, r *http.Request, w http.ResponseWriter) {
	fmt.Println("\n\n----------------------------------------------\nGot request...")

	stripe.Key = cfg.StripeKey

	dump, err := httputil.DumpRequest(r, true)
	if err == nil {
		fmt.Printf("%v: %s\n\n", time.Now(), dump)
	}

	token := r.FormValue("stripeToken")
	email := r.FormValue("stripeEmail")

	if len(token) == 0 {
		http.Error(w, "Error 11: No payment has been deducted. Please contact the site owner to proceed. Mentioning the error code (11) will help.", http.StatusInternalServerError)
		fmt.Println("No token!")
		return
	}

	if len(email) == 0 {
		http.Error(w, "Error 98: No payment has been deducted. Please contact the site owner to proceed. Mentioning the error code (98) will help.", http.StatusInternalServerError)
		fmt.Println("No email!")
		return
	}

	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(cfg.Amount),
		Currency:    stripe.String(string(stripe.CurrencyUSD)),
		Description: stripe.String(cfg.ChargeDescription),
	}
	params.SetSource(token)

	fmt.Println("Processing payment for " + email)

	charge, err := charge.New(params)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error 101 - No payment has been deducted. Please contact the site owner to proceed. Mentioning the error code (101) will help.", http.StatusInternalServerError)
		return
	}

	fmt.Println("Success!")

	b, err := json.MarshalIndent(charge, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(charge)
	} else {
		fmt.Println(string(b))
	}

	sendMail(cfg, email)

	fmt.Println("Redirecting...")
	http.Redirect(w, r, cfg.RedirectAfter, 303)
}

func sendMail(cfg config, recipient string) {
	fmt.Println("Sending email to " + recipient)

	mg := mailgun.NewMailgun(cfg.MailgunDomain, cfg.MailgunPvtKey)

	message := mg.NewMessage(cfg.MailFrom, cfg.MailSubject, cfg.MailTxt, recipient)

	message.SetHtml(cfg.MailHtml)
	message.AddInline("click.png")

	resp, id, err := mg.Send(message)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Failed sending Email!")
		return
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	fmt.Println("Email sent to " + recipient)
}

func validateConfigx1(cfg config) (config, error) {
	if cfg.Amount == 0 {
		return cfg, errors.New("Amount missing in configuration")
	}
	if isEmpty(cfg.ChargeDescription) {
		return cfg, errors.New("ChargeDescription missing in configuration")
	}
	if isEmpty(cfg.RedirectAfter) {
		return cfg, errors.New("RedirectAfter missing in configuration")
	}
	if isEmpty(cfg.MailFrom) {
		return cfg, errors.New("MailFrom missing in configuration")
	}
	if isEmpty(cfg.MailSubject) {
		return cfg, errors.New("MailSubject missing in configuration")
	}
	if isEmpty(cfg.MailTxt) {
		return cfg, errors.New("MailTxt missing!")
	}
	if isEmpty(cfg.MailHtml) {
		return cfg, errors.New("MailHtml missing!")
	}
	if isEmpty(cfg.StripeKey) {
		return cfg, errors.New("StripeKey missing in configuration")
	}
	if isEmpty(cfg.MailgunDomain) {
		return cfg, errors.New("MailgunDomain missing in configuration")
	}
	if isEmpty(cfg.MailgunPvtKey) {
		return cfg, errors.New("MailgunPvtKey missing in configuration")
	}
	return cfg, nil
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

type config struct {
	Amount            int64
	ChargeDescription string
	RedirectAfter     string

	MailFrom    string
	MailSubject string

	MailTxt  string
	MailHtml string

	StripeKey string

	MailgunDomain string
	MailgunPvtKey string
}
