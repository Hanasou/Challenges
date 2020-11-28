using System;
using System.Linq;
using System.Collections.Generic;

namespace Challenges
{
    public class TreeNode {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int val=0, TreeNode left=null, TreeNode right=null) {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }
    class Program
    {
        /**
        Given a binary tree root, a node X in the tree is "good" if in the path from root to X,
        there are no nodes with a value greater than X.
        Return the number of good nodes in the binary tree.
        **/
        public int GoodNodes(TreeNode root) {
            // The brute force solution would be something like checking every path from
            // root to X for all X in the tree.
            int goodNodes = 0;

            void nodeRecurse(TreeNode root, int prev) {
                // We've reached the end of the tree.
                if (root == null) {
                    return;
                }
                int x = 0;
                // Check if the current node's value is >= the previous node's value
                if (root.val >= prev) {
                    // If it is, then we found a good node
                    // Note that we don't have to check all of them. If it's greater than
                    // the previous node's value, it's greater than all of them.
                    goodNodes++;
                    x = root.val;
                } else {
                    x = prev;
                }
                nodeRecurse(root.left, x);
                nodeRecurse(root.right, x);
            }
            nodeRecurse(root, root.val);
            return goodNodes;
        }

        /**
        Given a matrix that consists of 0s and 1s, find the distance of the nearest 0
        for each cell. The distance between two adjacent cells is 1.
        **/
        public int[][] UpdateMatrix(int[][] matrix) {
            int rowLen = matrix.Length;
            int colLen = matrix[0].Length;
            int[][] distanceMatrix = new int[rowLen][];
            // fill the array with 0s
            for (int i = 0; i < rowLen; i++) {
                for (int j = 0; j < colLen; j++) {
                    distanceMatrix[i][j] = Int32.MaxValue;
                }
            }

            // Iterate over the input array
            for (int i = 0; i < rowLen; i++) {
                for (int j = 0; j < colLen; j++) {
                    if (matrix[i][j] == 0) {
                        distanceMatrix[i][j] = 0;
                    } else {
                        int north;
                        int south;
                        int west;
                        int east;
                        // Look at each adjacent cell
                        var adjacentCells = new List<int>();
                        if (i > 0) {
                            north = matrix[i-1][j];
                            adjacentCells.Add(north);
                        }
                        if (i < rowLen - 1) {
                            south = matrix[i+1][j];
                            adjacentCells.Add(south);
                        }
                        if (j > 0) {
                            west = matrix[i][j-1];
                            adjacentCells.Add(west);
                        }
                        if (j < colLen) {
                            east = matrix[i][j+1];
                            adjacentCells.Add(east);
                        }
                        if (adjacentCells.Contains(0)) {
                            distanceMatrix[i][j]++;
                        }
                    }
                }
            }
            return distanceMatrix;
        }

        /**
        This is very weird and complicated
        We have int arrays A and B, and we want to draw connecting lines between them
        A connecting line is a number from A[i] and B[j] such that
        A[i] == B[j] and the line we draw doesn't intersect with another line
        A: 1 4 2
        B: 1 2 4
        Find the number of connecting lines we can draw
        In this case, we can draw 2. 1 to 1, and 4 to 4, or 2 to 2.
        **/
        public int MaxCrossedLines(int[] A, int[] B) {
            // Let's define how we get a connecting line
            // Any line where elements are parallel is included in our count
            // So if A[i] == B[j] and i == j, then it is included.
            // Intuitively, the further apart i and j are, the more conflicts we can come across.
            // How do we know if two lines intersect each other?
            // If we have a line where i > j, and another line where i < j.
            /**
            2, 5,1,2,5
            10,5,2,1,5,2
            Since we're going to be connecting 5 and 5, that means we can't have any of the previous elements
            getting connected to the later elements.
            **/
            return 0;
        }

        /**
        There are 3n piles of coins of varying size, you and your friends will take piles of coins as follows
        In each step, you will choose any 3 piles of coins
        Of your choice, Alice will pick the pile with the max number of coins
        You will pick the next pile with the max number of coins
        Bob will pick the last pile
        Repeat until there are no more piles
        Given an int[] piles where piles[i] is the number of coins in the ith pile
        Return the max number of coins you can get
        **/
        public static int MaxCoins(int[] piles) {
            // Basically, each round, we choose three piles of coins.
            // Alice will pick the largest pile
            // I will pick the next largest pile.
            // Bob will get the last one
            // Basically, I have to pick triplets where the 2nd greatest pile in each of them
            // adds up to the largest possible value
            // For the last pile we take we should ostensibly take the lowest one

            var pilesList = piles.ToList<int>();
            var sortedList = pilesList.OrderByDescending(i => i).ToList();
            int rounds = piles.Length / 3;
            int myCoins = 0;
            for (int i = 1; i <= rounds * 2; i+=2) {
                // Grab the highest, second highest, and lowest values in piles
                int myPile = sortedList[i];
                myCoins += myPile;
            }
            return myCoins;
        }

        /**
        Given an int[] nums, find the maximum possible sum of elements of the array such that
        it is divisible by three.
        [3,6,5,1,8]
        Answer is 18. Pick numbers 3, 6, 1, and 8 to get the max number divisible by three.
        **/
        public static int MaxSumDivThree(int[] nums) {
            // Start with the largest number first
            // See what you can add to it to make a pair that's divisible by 3
            var numList = nums.ToList<int>();
            var sortedList = numList.OrderByDescending(i => i).ToList();
            var chosenNumbers = new List<int>();
            int j = sortedList.Count - 1;
            for(int i = 0; i < j; i++) {
                if ((sortedList[i] + sortedList[j]) % 3 == 0) {
                    chosenNumbers.Add(sortedList[i]);
                    chosenNumbers.Add(sortedList[j]);
                    sortedList.Remove(sortedList[j]);
                    j--;
                }
            }
            return chosenNumbers.Sum();        
        }

        /**
        Given an integer array nums, return the number of longest increasing subsequences.
        nums = [1,3,5,4,7]
        Answer is 2 because the length of the longest increasing subsequence is 4.
        There are two increasing subsequences of length 4,
        1,3,4,7 and 1,3,5,7
        **/
        public static int FindNumberOfLIS(int[] nums) {
            return 0;
        }
        static void Main(string[] args)
        {
            Console.WriteLine(MaxSumDivThree(new int[] {1,2,3,4,4}));
        }
    }
}
