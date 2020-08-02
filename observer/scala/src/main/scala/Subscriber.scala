package observer

class Subscriber(val publisher: Publisher) {
  var subscription: Int = _

  def shutdown(): Unit = {
    println("Subscriber: Unsubscribing...")
    publisher.unsubscribe(subscription)
  }

  def speak(d: String): Unit = println(s"Subscriber: Publisher says ${'"'}$d${'"'}")
}
