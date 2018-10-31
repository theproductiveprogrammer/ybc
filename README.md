YBC - A Simple Stripe Backend
=============================

![YBC](payment.png)

[Stripe](https://stripe.com/) is simple, brilliant solution for small
businesses wanting to accept payments from their users.

YBC is a simple backend for bloggers who have some access to a server
and want to host their own payment service. Deploying YBC allows them to
sell their products and get an email sent out to their customers with
the product links on success.

## Requirements
1. Access to a Server where they can install
   [`Golang`](https://golang.org)
2. A `Stripe` [developer account](https://stripe.com/docs)
3. A `Mailgun` [developer account](https://www.mailgun.com/email-api)
3. An [`nginx`](https://nginx.org) SSL-Offloading front end

(**To Ponder**: Ok this seems too complicated! Can this be simplified so
it's more usable without having to go for a managed service?)


## How it Works
All that is required is:
* the price of the product
* the mail template to send across to the user

And YBC will do the rest. On the front-end, we use
[`Checkout`](https://stripe.com/docs/checkout) which is a beautifully
designed, extensively UX-tested experience designed by
[Stripe](https://stripe.com).


## Setup

Please copy the code from `buy.html` and **MAKE SURE YOU CHANGE THE
STRIPE KEY** to your own key.

Then create a `ybc.cfg` file that contains the following:

```txt
Amount = 4800
ChargeDescription = "YBC-Acne Healing Diet"
RedirectAfter = "https://www.yourbeautychronicles.com/thank-you-page/"

MailFrom = "Anjali Lobo <yourbeautychronicles@gmail.com>"
MailSubject = "The Acne Healing Diet"

StripeKey = "<<YOUR STRIPE KEY>>"

MailgunDomain = "<<YOUR MAILGUN DOMAIN>>"
MailgunPvtKey = "<<YOUR MAILGUN PRIVATE KEY>>"
```

Also create two files: `mail.txt` and `mail.html` that contain your
welcome mail (in text and HTML format).

Once this is done, you are ready to go!

## Feedback/Issues/Suggestions
[All inputs welcome](https://github.com/theproductiveprogrammer/ybc/issues)
