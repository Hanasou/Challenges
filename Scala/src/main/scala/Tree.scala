import scala.collection.mutable.ArrayBuffer

object Tree {
  class BinaryTree(var value: Int = 0, var left: BinaryTree = null, var right: BinaryTree = null) {

    def insertLeft(node: Int): Unit = {
      if (this.left == null) {
        this.left = new BinaryTree(node)
      } else {
        // Replace left subtree with new BinaryTree
        val t = new BinaryTree(node) // Make new subtree
        // Set left subtree of new node to current subtree
        t.left = this.left
        // Set new subtree to left of this tree
        this.left = t
      }
    }

    // Kinda the same thing as insertLeft
    def insertRight(node: Int): Unit = {
      if (this.right == null) {
        this.right = new BinaryTree(node)
      } else {
        val t = new BinaryTree(node)
        t.right = this.right
        this.right = t
      }
    }
  }

  // Given a Binary Tree, determine if it's a BST
  // A Binary Tree is a BST if every value of its left subtree is less than itself
  // and every value of its right subtree is greater than itself
  def verifyBST(tree: BinaryTree): Boolean = {
    val minValue = Int.MinValue
    val maxValue = Int.MaxValue
    def helper(tree: BinaryTree, lower: Int, upper: Int): Boolean = {
      if (tree == null) {
        return true
      }

      val value = tree.value
      if (value <= lower || value >= upper) {
        return false
      }

      // Compare right subtree with root as lower bound
      if (!helper(tree.right, value, upper)) {
        return false
      }
      // Compare left subtree with root as upper bound
      if (!helper(tree.left, lower, value)) {
        return false
      }
      true
    }
    helper(tree, minValue, maxValue)
  }

  // Basically do breadth first search and print I guess
  def treeLevelPrint(tree: BinaryTree): Unit = {

  }

  // Take out all the nodes whose values are outside of min and max
  def trimTree(tree: BinaryTree, min: Int, max: Int): BinaryTree = {
    null
  }
  def main(args: Array[String]): Unit = {
    val array = Array[Int](0)
    println(array.mkString(" "))
  }
}
