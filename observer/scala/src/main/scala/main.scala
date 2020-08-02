package observer

object main extends App {
  val p = new Publisher()
  val s = new Subscriber(p)

  // Subscribe to publisher with subscriber's Speak func
  p.subscribe(s.speak)

  // Start publisher's time wasting
  p.start()

  // Unsubscribe subscriber from publisher
  s.shutdown()
}
