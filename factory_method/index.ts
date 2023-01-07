interface Payment {
    doPayment(): void;
}

class CardPayment implements Payment {
    doPayment(): void {
        console.log("Doing credit card payment");
    }
}

class TransferPayment implements Payment {
    doPayment(): void {
        console.log("Doing 3.0 transfer payment");
    }
}

abstract class DoPayment {
    // factory method
    abstract getPaymentType(): Payment;

    pay(): void {
        const payment = this.getPaymentType();
        payment.doPayment();
    }
}

class DoPaymentCard extends DoPayment {
    getPaymentType(): Payment {
        return new CardPayment();
    }
}

class DoPaymentTransfer extends DoPayment {
    getPaymentType(): Payment {
        return new TransferPayment();
    }
}

function getPaymentMethod(method: string): DoPayment {
  let doPayment: DoPayment;

  switch (method) {
    case "transfer": {
      doPayment = new DoPaymentTransfer();
      break;
    }
    case "card": {
      doPayment = new DoPaymentCard();
      break;
    }
    default: {
      console.log("Payment method not allowed");
      process.exit();
    }
  }

  return doPayment;
}

const PAYMENT_METHOD = {
  CARD: "card",
  TRANSFER: "transfer",
};

let paymentType = PAYMENT_METHOD.TRANSFER;
const doPayment = getPaymentMethod(paymentType);
doPayment.pay();
