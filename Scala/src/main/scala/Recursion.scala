import scala.collection.mutable.ArrayBuffer

object Recursion {

  def factorial(n: Int): Int = {
    if (n == 0) {
      return 1
    }
    n * factorial(n - 1)
  }

  /*
  Add up all the digits of an Integer.
  Do it *recursively*
   */
  def sumDigits(n: Int): Int = {
    if (n == 0) {
      return 0
    }
    val lastDigit = n % 10
    val nextDigit = n / 10
    lastDigit + sumDigits(nextDigit)
  }

  def reverseString(s: String): String = {
    def helper(original: String, reversed: String, place: Int): String = {
      if (place < 0) {
        return reversed
      }
      val newReversed = reversed + original.charAt(place)
      val newPlace = place - 1
      helper(original, newReversed, newPlace)
    }
    helper(s, "", s.length() - 1)
  }

  /*
  Take in a string and return a list of all possible permutations
   */
  def stringPermutations(s: String): ArrayBuffer[String]= {
    // Initialize output
    val output = ArrayBuffer[String]()

    // A string with length of one has a single permutation: itself
    if (s.length() == 1) {
      return ArrayBuffer(s)
    }

    // Loop through all the characters in the string
    for((c,i) <- s.zipWithIndex) {
      // Do permutations on all substrings excluding the character we're examining
      // Loop through all those permutations
      for (perm <- stringPermutations(s.substring(0,i) + s.substring(i + 1))) {
        // Insert the character that we're examining to the permutation
        output += c + perm
      }
    }

    // Return output
    output
  }

  def fibonacci(n: Int): Int = {
    val memory = new Array[Int](n + 1)
    def fibDp(n: Int, mem: Array[Int]): Int = {
      if (n == 0 || n == 1) {
        return n
      }
      if (mem(n) == 0) {
        mem(n) = fibDp(n - 1, mem) + fibDp(n - 2, mem)
      }

      mem(n)
    }
    fibDp(n, memory)
  }

  /*
  Given n and a list of distinct coins, give the fewest number of coins required to get n
  e.g. n = 10, coinList = [1,5,10]
  return 1
   */
  def coinChange(n: Int, coinList: Array[Int]): Int = {
    val mem = new Array[Int](n + 1)

    def helper(n: Int, coinList: Array[Int], mem: Array[Int]): Int = {

      // Recursive function to find the minimum
      // Max possible value is n
      var minCoins = n

      // Base Case. We're going to be subtracting n and recursively calling this function.
      // Eventually we'll get an n that matches a coin
      if (coinList contains n) {
        mem(n) = 1
        return 1
      }
      // We're using values in mem to calculate other values.
      // Instead of recalculating, we'll pick it up from mem.
      if (mem(n) > 0) {
        return mem(n)
      }

      // Recursively calculate minCoins
      // Loop through all the coins that are <= n
      val validCoins = coinList.filter(_ <= n)
      for (c <- validCoins) {
        // numCoins represents number of coins that can go to n
        // e.g. n = 10, c = 1. numCoins would be 1 + numCoins when n = 9
        val numCoins = 1 + helper(n - c, coinList, mem)

        // Modify min if numCoins is less than it
        if (numCoins < minCoins) {
          minCoins = numCoins
          mem(n) = minCoins
        }
      }

      minCoins
    }
    0
  }


  def main(args: Array[String]): Unit = {
    val mem = Array[Int](1,5,6,7,2,3)
    val filtered = mem.filter(_ > 5)
    println(filtered.mkString(" "))
  }
}