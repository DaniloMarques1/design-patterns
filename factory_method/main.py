from abc import abstractclassmethod

class Shipping:
    @abstractclassmethod
    def deliver(self):
        pass

class Truck(Shipping):
    def deliver(self):
        print("deliverying by truck")

class Ship(Shipping):
    def deliver(self):
        print("deliverying by ship")

class ShippingDelivery:
    @abstractclassmethod
    def createTransport(self):
        pass

    def deliver(self):
        shipping = self.createTransport()
        shipping.deliver()

class ShippingByTruck(ShippingDelivery):
    def createTransport(self):
        return Truck()

class ShippingByShip(ShippingDelivery):
    def createTransport(self):
        return Ship()

def main():
    shippingType = "truck"
    shipping = None # in theory this is of the type ShippingDelivery
    if (shippingType == "truck"):
        shipping = ShippingByTruck()
    elif (shippingType == "ship"):
        shipping = ShippingByShip()
    else:
        print("wrong shipping type")
        exit()

    shipping.deliver()

main()
