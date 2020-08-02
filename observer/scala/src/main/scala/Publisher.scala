package observer

class Publisher {
  var subscriptions: List[String => Unit] = List()

  def start(): Unit = {
    println("Publisher: Wasting time...")
    Thread.sleep(5000)
    end()
  }

  def end(): Unit = {
    println("Publisher: Notifying subscribers that I am done wasting time...")
    notify("Done")
  }

  def notify(d: String): Unit = {
    for (s <- subscriptions) s(d)
  }

  def subscribe(f: String => Unit): Int = {
    println("Publisher: Subscribing a subscriber...")
    subscriptions = subscriptions ++ List(f)
    subscriptions.length - 1
  }

  def unsubscribe(i: Int): Unit = {
    println("Publisher: Unsubscribing a subscriber...")
    subscriptions = subscriptions.patch(i, Nil, 1)
  }
}
