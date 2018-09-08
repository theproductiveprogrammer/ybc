YBC - A Simple Stripe Backend
=============================

![YBC](ybc.png)

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

```txt
[keys]
stripe=<stripe key>
mailgun=<mailgun key>

[myproduct1]
price=4800  # $48 in cents
```
```txt
Hi!

Thanks for buying my product.
Click <a href=/some/product.delivery?user={{email}}>here</a> for your copy.

Thanks!
```

And YBC will do the rest. On the front-end, we use
[`Checkout`](https://stripe.com/docs/checkout) which is a beautifully
designed, extensively UX-tested experience designed by
[Stripe](https://stripe.com). To do this, simply update the following
parameters in the code and add it to your site.

```html
<form action="your-server-side-code" method="POST">
  <script
    src="https://checkout.stripe.com/checkout.js" class="stripe-button"
    data-key="pk_test_TYooMQauvdEDq54NiTphI7jx"
    data-amount="999"
    data-name="Stripe.com"
    data-description="Widget"
    data-image="https://stripe.com/img/documentation/checkout/marketplace.png"
    data-locale="auto"
    data-zip-code="true">
  </script>
</form>
```

**NB**: that you need to type in the `amount` being charged here to show it
to your customer when the form comes up.

## Feedback/Issues/Suggestions
[All inputs welcome](https://github.com/theproductiveprogrammer/ybc/issues)
