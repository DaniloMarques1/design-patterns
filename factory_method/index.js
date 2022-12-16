var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
var CardPayment = /** @class */ (function () {
    function CardPayment() {
    }
    CardPayment.prototype.doPayment = function () {
        console.log("Doing credit card payment");
    };
    return CardPayment;
}());
var TransferPayment = /** @class */ (function () {
    function TransferPayment() {
    }
    TransferPayment.prototype.doPayment = function () {
        console.log("Doing 3.0 transfer payment");
    };
    return TransferPayment;
}());
var DoPayment = /** @class */ (function () {
    function DoPayment() {
    }
    DoPayment.prototype.pay = function () {
        var payment = this.getPaymentType();
        payment.doPayment();
    };
    return DoPayment;
}());
var DoPaymentCard = /** @class */ (function (_super) {
    __extends(DoPaymentCard, _super);
    function DoPaymentCard() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    DoPaymentCard.prototype.getPaymentType = function () {
        return new CardPayment();
    };
    return DoPaymentCard;
}(DoPayment));
var DoPaymentTransfer = /** @class */ (function (_super) {
    __extends(DoPaymentTransfer, _super);
    function DoPaymentTransfer() {
        return _super !== null && _super.apply(this, arguments) || this;
    }
    DoPaymentTransfer.prototype.getPaymentType = function () {
        return new TransferPayment();
    };
    return DoPaymentTransfer;
}(DoPayment));
var paymentType = "transfer";
var doPayment;
switch (paymentType) {
    case "transfer":
        doPayment = new DoPaymentTransfer();
        break;
    case "card":
        doPayment = new DoPaymentCard();
        break;
    default:
        console.log("Payment method not allowed");
        process.exit();
}
doPayment.pay();
