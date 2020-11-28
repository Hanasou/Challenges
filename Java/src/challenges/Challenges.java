package challenges;

import java.util.*;

import static challenges.ListNode.fromList;
import static challenges.ListNode.listToString;

public class Challenges {

    /**
     * Return the number of longest increasing subsequences for int[] nums
     * Input: [1,3,5,4,7]
     * Output: 2. [1,3,5,7] and [1,3,4,7] are the two longest increasing subsequences
     * @param nums
     * @return
     */
    public static int findNumberOfLIS(int[] nums) {
        // First of all, what's a subsequence?
        // I assume it's any sequence of numbers in the array
        // A sequence appears to be that each number in the sequence must come after
        // the previous number in the original array.
        // [1,3,4,5] would not be valid because 4 comes after 5 in the original array
        // Increasing means that the next number is greater than the previous one

        // First, let's try finding a single increasing subsequence
        // This finds one of the longest increasing subsequences in the list
        // How do we find all of them?
        List<Integer> subsequence = new ArrayList<>();
        int prev = Integer.MIN_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int currNumber = nums[i];
            if (i != 0) {
                prev = nums[i - 1];
            }
            if (currNumber > prev) {
                subsequence.add(currNumber);
            }
        }
        System.out.println(subsequence.toString());
        return 0;
    }

    /**
     * Given the head of a linked list and int n, remove the nth node from the end
     * of the list and return its head.
     * @param head
     * @param n
     * @return
     */
    public static ListNode removeNthFromEnd(ListNode head, int n) {
        // See if you can do this in one pass
        // It's kind of difficult since we don't know the size of the linked list
        // We can do this in one pass by using two pointers
        // The first pointer goes to n

        // Initialize a dummy node to make resolving some corner cases easier
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        // Initialize pointers
        ListNode first = dummy;
        ListNode second = dummy;

        // Move up the first pointer so that the first and second pointers are n nodes apart
        for (int i = 1; i <= n + 1; i++) {
            first = first.next;
        }
        // Move first up to the end
        // Once first reaches the end, we know that second is at the (n-1)th node
        while (first != null) {
            first = first.next;
            second = second.next;
        }
        // We just remove the node by skipping over it
        second.next = second.next.next;
        // the original head never gets modified, so we return dummy.next instead of head
        return dummy.next;
    }

    /**
     * Given a string array of equations, each string equations[i] has length 4
     * and takes the form "a==b" or "a!=b". a and b are lowercase letters (not necessarily different)
     * and represent one-letter variable names
     * Return true only if it's possible to assign integers to variable names to satisfy all given equations
     * @param equations
     * @return
     */
    public boolean equationsPossible(String[] equations) {
        // First let's try iterating over all the equations
        for (int i = 0; i < equations.length; i++) {
            // Let's try parsing the equation
            // Or maybe this problem is just way too bizarre for me to even try
        }
        return false;
    }

    /**
     * Merge two sorted lists. Not much else to say I guess
     * @param l1
     * @param l2
     * @return
     */
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        // Making a dummy makes corner cases easier
        ListNode dummy = new ListNode(0);
        // We're going to be continuously changing the pointer of dummy
        // Head is always going to point to the head of dummy
        ListNode head = dummy;
        while(l1 != null && l2 != null) {
            // check the values of l1 and l2
            // insert the one that's lower into dummy
            if (l1.val >= l2.val) {
                dummy.next = l2;
                l2 = l2.next;
            } else {
                dummy.next = l1;
                l1 = l1.next;
            }
            dummy = dummy.next;
        }
        // Fill dummy with the leftover nodes from whichever list has them.
        if (l1 != null) {
            dummy.next = l1;
        } else {
            dummy.next = l2;
        }
        // Remember that head is still at the starting position while
        // dummy was getting changed.
        return head.next;
    }

    /**
     * int[] nums is sorted in ascending order
     * nums is rotated at some pivot that is unknown to us
     * (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
     * If target is found, return its index. Otherwise, return -1.
     * @param nums
     * @param target
     * @return
     */
    public int search(int[] nums, int target) {
        // For a search algorithm normally we would do binary search
        // How would be modify binary search so that it would work for this?
        // After checking the middle, we can try checking the other
        // two halves before actually setting it somewhere
        int start = 0;
        int end = nums.length - 1;
        while (start <= end) {
            // assign middle because start and end will change
            int middle = (end + start) / 2;
            int val = nums[middle];
            int rightVal = nums[end];
            if (val == target) {
                return middle;
            }

            // Check if middle val is less than right val
            // If this is true, that means the pivot is on the left
            if (val < rightVal) {
                // We know the pivot is on the left
                // Therefore, we have to be sure that our target is both greater than val
                // and less than or equal to rightVal to be sure if our target is on the right
                if (target > val && target <= rightVal) {
                    start = middle+1;
                } else {
                    end = middle-1;
                }
            }
            // Entering this block means that val is >= rightVal
            // That means the pivot is somewhere on the right
            else {
                // We know the pivot is on the right
                // That means we know that there are values here that are greater than val
                // and there are values here that are less than rightVal
                // If our target meets either of these two criteria, we know it's on the right
                if (target > val || target <= rightVal) {
                    start = middle+1;
                } else {
                    end = middle-1;
                }
            }
        }
        return -1;
    }

    /**
     * Given an array of integers nums sorted in ascending order, find the starting and ending position
     * of a given target value.
     * If target is not found in the array, return [-1, -1].
     * Input: nums = [5,7,7,8,8,10], target = 8
     * Output: [3,4]
     * @param nums
     * @param target
     * @return
     */
    public static int[] searchRange(int[] nums, int target) {
        // Try doing binary search and then creeping along the edges to see how many elements there are
        // when we actually find the target.
        int left = 0;
        int right = nums.length - 1;
        while (left <= right) {
            int mid = (right + left) / 2;
            int val = nums[mid];
            if (val == target) {
                int[] range = {mid, mid};
                int p1 = mid;
                int p2 = mid;
                // Let's just have two while loops for p1 and p2
                while (p1 >= 0 && nums[p1] == val) {
                    range[0] = p1;
                    p1--;
                }
                while (p2 < nums.length && nums[p2] == val) {
                    range[1] = p2;
                    p2++;
                }
                return range;
            }
            if (val > target) {
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        return new int[] {-1,-1};
    }

    /**
     * Given an array of distinct integers candidates and a target integer target,
     * return a list of all unique combinations of candidates where the chosen numbers sum to target.
     * You may return the combinations in any order.
     * The same number may be chosen from candidates an unlimited number of times.
     * Two combinations are unique if the frequency of at least one of the chosen numbers is different.
     * Input: candidates = [2,3,5], target = 8
     * Output: [[2,2,2,2],[2,3,3],[3,5]]
     * Note that two is used four times in the first combination. Two can be used again in the second
     * combination, but it can't be used four times. The same applies for all other numbers
     * @param candidates
     * @param target
     * @return
     */
    public static List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> combinations = new ArrayList<>();
        List<Integer> current = new ArrayList<>();
        getSum(candidates, target, 0, current, combinations);
        return combinations;
    }

    public static void getSum(int[] candidates, int target, int index,
                              List<Integer> current, List<List<Integer>> result) {
        // We're going to be continuously subtracting from target
        if (target == 0) {
            result.add(new ArrayList<>(current));
            return;
        }
        // Loop through candidates
        for (int i = index; i < candidates.length; i++) {
            int val = candidates[i];
            if (target - val >= 0) {
                // The current value is part of the combination
                current.add(val);
                // We do this again. Update the target and index.
                // We have to start at a new index so we don't recalculate results
                getSum(candidates, target - val, i, current, result);
                // Not entirely sure why we do this
                // I guess we're cleaning the array so then we can reuse it
                // After we calculate one of the combination sums?
                current.remove(current.size() - 1);
            }
        }
    }

    /**
     * We're going to practice dp once again by redoing fib
     * I've done this quite a lot let's see if I can remember
     * Nope I still don't remember from scratch
     * Return the nth number in the fibonacci sequence
     * @param n
     * @return
     */
    public static int fibNumbers(int n) {
        int[] mem = new int[n];
        return fibDp(n, mem);
    }
    public static int fibDp(int n, int[] mem) {
        // Trivial case
        if (n == 0 || n == 1) {
            return n;
        }
        // We've gotten to a spot in our memory that we haven't calculated yet
        if (mem[n] == 0) {
            mem[n] = fibDp(n - 1, mem) + fibDp(n - 2, mem);
        }
        return mem[n];
    }

    /**
     * Given a collection of distinct integers, return all possible permutations.
     * @param nums
     * @return
     */
    public List<List<Integer>> permute(int[] nums) {
        // I probably have to use some sort of backtracking algorithm for this
        // Something along the lines of iterating over nums
        // For each element, add the other two
        List<List<Integer>> result = new ArrayList<>();
        permuteHelper(nums, result,  new ArrayList<Integer>());
        return result;
    }

    public void permuteHelper(int[] nums, List<List<Integer>> result, List<Integer> current) {
        // Trivial case
        // If we know that current is a permutation, add it into result
        int numsLength = nums.length;
        if (current.size() == numsLength) {
            result.add(new ArrayList<Integer>(current));
            return;
        }
        for (int i = 0; i < numsLength; i++) {
            int val = nums[i];
            // Find some condition to add this result into our current permutation
            // Well first of all, we don't add it if it's already in
            if (current.contains(val)) {
                continue;
            }
            current.add(val);
            permuteHelper(nums, result, current);
            current.remove(current.size() - 1);
        }
    }

    /**
     * Given an nxn 2d matrix, rotate it 90 degrees in place
     * Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
     * Output: [[7,4,1],[8,5,2],[9,6,3]]
     * @param matrix
     */
    public static void rotate(int[][] matrix) {
        // What appears to be the case is that we can take all the elements from each array
        // And put them each into different arrays at the last - nth element
        // For example, the first array contains [1,2,3], that means to rotate the matrix,
        // 1 belongs to the last element of the first array, 2 belongs to the last element
        // of the second array, 3 belongs to the last element of the third array.
        // For the second array, [4,5,6], we do the same thing except each element belongs to the
        // second last element of each array.
        // We have to figure out a way to do this in place as well. This might mean swapping elements.

        // Let's try to figure out how we can use the indexes to see which element belongs in which array
        // The index of matrix[row][0] belongs to the array of the 0th row.
        // The index of matrix[row][1] belongs to the array of the 1st row and so on.
        // The first thing we should do is try to swap elements so that each element belongs to the correct array
        // We can reorder each individual one later.
        for (int i = 0; i < matrix.length; i++) {
            for (int j = i; j < matrix[0].length; j++) {
                int val = matrix[i][j];
                // We want val to be in the correct array
                if (i == j) {
                    // We don't swap here because his means that val is in the right array
                    continue;
                } else {
                    int swapWith = matrix[j][i];
                    matrix[j][i] = val;
                    matrix[i][j] = swapWith;
                }
            }
        }
        // Reverse each array
        for (int i = 0; i < matrix.length; i++) {
            int p1 = 0;
            int p2 = matrix[i].length - 1;
            while (p1 < p2) {
                int left = matrix[i][p1];
                int right = matrix[i][p2];
                matrix[i][p1] = right;
                matrix[i][p2] = left;
                p1++;
                p2--;
            }
        }
    }

    /**
     * Return the longest palindromic substring
     * @param s
     * @return
     */
    public String longestPalindrome(String s) {
        // First check if string is null
        if (s == null || s.length() < 1) {
            return "";
        }
        int start = 0, end = 0;
        // iterate over each character in the string
        for (int i = 0; i < s.length(); i++) {
            // There are two cases we have to concern ourself with
            // First case is an odd numbered palindrome, where every character
            // around the middle character is the same and the middle character is whatever
            int len1 = expandAroundCenter(s, i, i);
            // The second case is an even numbered palindrome, where every character
            // around the center is equal to its opposing member
            int len2 = expandAroundCenter(s, i, i+1);
            // Get length of the longest of these two palindromes
            int len = Math.max(len1, len2);
            // Check if the length of this current palindrome is greater than the length of the previous one
            // Keep in mind that end and start are tracking the position of the currently longest palindrome
            if (len > end - start) {
                // These two formulas give us the indices of our longest palindrome
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        // Return the longest palindromic substring
        return s.substring(start, end + 1);
    }
    private int expandAroundCenter(String s, int left, int right) {
        // Is there a particular reason we're assigning our parameters to
        // new variables instead of just using them directly?
        int L = left, R = right;
        while (L >= 0 && R < s.length() && s.charAt(L) == s.charAt(R)) {
            L--;
            R++;
        }
        return R - L - 1;
    }

    /**
     * Given an array of non-negative integers, you are initially positioned at the first index of the array.
     * Each element in the array represents your maximum jump length at that position.
     * Determine if you are able to reach the last index.
     * Input: nums = [2,3,1,1,4]
     * Output: true
     * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
     * @param nums
     * @return
     */
    public boolean canJump(int[] nums) {
        // Probably do some backtracking stuff here
        // Maybe also try a DP approach
        return jumpHelper(nums, 0);
    }
    private static boolean jumpHelper(int[] nums, int index) {
        if (index == nums.length - 1) {
            return true;
        }
        // Make sure that we don't go past the range of nums
        int furthestJump = Math.min(index + nums[index], nums.length - 1);
        // Try all sorts of jumps for each number
        for (int i = index; i < furthestJump; i++) {
            if (jumpHelper(nums, i + 1)) {
                return true;
            }
        }
        return false;
    }

    // GOOD means that we are able to jump to the last index from this position
    // BAD means that we are not able to
    // UNKNOWN means that we don't know if we can
    enum Index {
        GOOD, BAD, UNKNOWN;
    }

    public boolean canJumpDp(int[] nums) {
        // Initialize the enum array as UNKNOWNS
        // The last element would be GOOD since that's what we're aiming for
        Index[] mem = new Index[nums.length];
        Arrays.fill(mem, Index.UNKNOWN);
        mem[mem.length - 1] = Index.GOOD;
        return jumpHelperDp(nums, 0, mem);
    }
    private static boolean jumpHelperDp(int[] nums, int index, Index[] mem) {
        // Check if it's good or bad
        // Return true for GOOD, false for BAD.
        if (mem[index] != Index.UNKNOWN) {
            return mem[index] == Index.GOOD;
        }
        int furthestJump = Math.min(index + nums[index], nums.length - 1);
        for (int i = index; i < furthestJump; i++) {
            if (jumpHelperDp(nums, i + 1, mem)) {
                mem[index] = Index.GOOD;
                return true;
            }
        }
        mem[index] = Index.BAD;
        return false;
    }

    /**
     * You are given coins of different denominations and a total amount of money amount.
     * Write a function to compute the fewest number of coins that you need to make up that amount.
     * If that amount of money cannot be made up by any combination of the coins, return -1.
     */
    public int coinChange(int[] coins, int amount) {
        if (amount < 1) {
            return 0;
        }
        int[] mem = new int[amount];
        return coinChangeHelper(coins, amount, mem);
    }

    private int coinChangeHelper(int[] coins, int remaining, int[] mem) {
        // Base cases
        // It takes 0 coins to get to 0.
        if (remaining == 0) {
            return 0;
        }
        // Overflow
        if (remaining < 0) {
            return -1;
        }
        // This means we're using a recalculated value which represents the min number of coins
        // it took to get to this spot. We can use this value to calculate what we want.
        if (mem[remaining - 1] != 0) {
            return mem[remaining - 1];
        }
        // Get the min value
        int min = Integer.MAX_VALUE;
        // We're going to be calling this function for each coin
        for (int coin : coins) {
            // Recursively calculate how many coins it takes to reach 0
            int result = coinChangeHelper(coins, remaining - coin, mem);
            // This is the current minimum number of coins it takes to reach a certain result
            if (result >= 0 && result < min) {
                min = 1 + result;
            }
        }
        // If we never modified the min, then we never replaced Integer.MAX_VALUE
        // so fill this spot with -1.
        mem[remaining - 1] = (min == Integer.MAX_VALUE) ? -1 : min;
        return mem[remaining - 1];
    }

    /**
     * Given an array of strings strs, group the anagrams together. You can return the answer in any order.
     * Input: strs = ["eat","tea","tan","ate","nat","bat"]
     * Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
     */
    public static List<List<String>> groupAnagrams(String[] strs) {
        // The first thing we can try to do is compare one string with every other string
        // Check if they're anagrams. If they are, put them in a new list.
        if (strs.length == 0) {
            return new ArrayList<>();
        }
        Map<String, List<String>> answers = new HashMap<>();
        for (String s : strs) {
            String key = getCharCounts(s).toString();
            if (!answers.containsKey(key)) {
                answers.put(key, new ArrayList<String>());
            }
            answers.get(key).add(s);
        }
        return new ArrayList(answers.values());
    }
    public static HashMap<Character, Integer> getCharCounts(String s) {
        HashMap<Character, Integer> counts = new HashMap<>();
        for (int i = 0; i < s.length(); i++) {
            Character c = s.charAt(i);
            if (!counts.containsKey(c)) {
                counts.put(c, 1);
            } else {
                counts.put(c, counts.get(c) + 1);
            }
        }
        return counts;
    }

    /**
     * A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
     * The robot can only move either down or right at any point in time.
     * The robot is trying to reach the bottom-right corner of the grid.
     * How many possible unique paths are there?
     * Input: m = 3, n = 7
     * Output: 28
     */
    public int uniquePaths(int m, int n) {
        return pathsRecurse(m, n, new int[m + 1][n + 1]);
    }
    public int pathsRecurse(int m, int n, int[][] mem) {
        // Out of bounds
        if (m == 0 || n == 0) {
            return 0;
        }
        // m == 1 && n == 1 represents the start.
        // We're basically starting at (m,n) and working our way backwards to (1,1)
        // If we get there, then that's a valid path.
        if (m == 1 && n == 1) {
            return 1;
        }
        // Each value on the grid represents how many different paths we can take to get there
        if (mem[m][n] != 0) {
            return mem[m][n];
        }
        // In order to calculate how many paths it takes to get to (m,n), find out how many
        // paths it takes in order to get to (m-1,n) and (m, n-1)
        return mem[m][n] = pathsRecurse(m - 1, n, mem) + pathsRecurse(m, n - 1, mem);
    }

    /**
     * Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right,
     * which minimizes the sum of all numbers along its path.
     */
    public int minPathSum(int[][] grid) {
        int[][] mem = new int[grid.length][grid[0].length];
        return pathRecurse(grid.length - 1, grid[0].length - 1, grid, mem);
    }
    public int pathRecurse(int row, int col, int[][] grid, int[][] mem) {
        if (row == 0 && col == 0) {
            return grid[row][col];
        }
        if (row == 0) {
            return mem[row][col] = pathRecurse(row, col - 1, grid, mem) + grid[row][col];
        }
        if (col == 0) {
            return mem[row][col] = pathRecurse(row - 1, col, grid, mem) + grid[row][col];
        }
        return mem[row][col] = grid[row][col] + Math.min(pathRecurse(row, col - 1, grid, mem),
                pathRecurse(row - 1, col, grid, mem));
    }

    /**
     * It takes n steps to reach the top
     * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
     */
    public int climbStairs(int n) {
        int[] mem = new int[n];
        return climbStairsDp(n, mem);
    }
    public int climbStairsDp(int n, int[] mem) {
        if (n < 0) {
            return 0;
        } else if (n == 0) {
            return 1;
        }
        if (mem[n] != 0) {
            return mem[n];
        }
        return mem[n] = climbStairsDp(n - 1, mem) + climbStairsDp(n - 2, mem);
    }

    /**
     * Given an array nums with n objects colored red, white, or blue,
     * sort them in-place so that objects of the same color are adjacent,
     * with the colors in the order red, white, and blue.
     * Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.
     */
    public static void sortColors(int[] nums) {
        // Try to solve this in one pass in place
        int redCount, whiteCount, blueCount;
        redCount = whiteCount = blueCount = 0;
        for (int i = 0; i < nums.length; i++) {
            int val = nums[i];
            if (val == 0) {
                redCount++;
            } else if (val == 1) {
                whiteCount++;
            } else if (val == 2) {
                blueCount++;
            }
        }
        for (int j = 0; j < nums.length; j++) {
            if (redCount > 0) {
                nums[j] = 0;
                redCount--;
            } else if (whiteCount > 0) {
                nums[j] = 1;
                whiteCount--;
            } else if (blueCount > 0) {
                nums[j] = 2;
                blueCount--;
            }
        }
    }

    public static List<List<Integer>> subsets(int[] nums) {
        Set<List<Integer>> sets = new HashSet<>();
        List<Integer> set = new ArrayList<>();
        for (int i = 0; i < nums.length; i++) {
            set.add(nums[i]);
        }
        subsetsHelper(sets, set);
        return new ArrayList<List<Integer>>(sets);
    }
    public static void subsetsHelper(Set<List<Integer>> sets, List<Integer> set) {
        if (set.isEmpty()) {
            sets.add(set);
        } else {
            sets.add(set);
            subsetsHelper(sets, set.subList(0, set.size() - 1));
        }
    }

    /**
     * Given an m x n board and a word, find if the word exists in the grid.
     *
     * The word can be constructed from letters of sequentially adjacent cells,
     * where "adjacent" cells are horizontally or vertically neighboring.
     * The same letter cell may not be used more than once.
     */
    public boolean exist(char[][] board, String word) {
        for (int i = 0; i < board.length; i++) {
            for (int j = 0; j < board[0].length; j++) {
                if (board[i][j] == word.charAt(0)) {

                }
            }
        }
        return false;
    }
    public boolean search(char[][] board, String word, int index, int i, int j) {
        // If we've found enough characters to make our word
        if (index == word.length()) {
            return true;
        }
        // If we're out of bounds
        if (i < 0 || j < 0 || i >= board.length || j >= board[0].length) {
            return false;
        }
        search(board, word, index + 1, i + 1, j);
        search(board, word, index + 1, i - 1, j);
        search(board, word, index + 1, i, j + 1);
        search(board, word, index + 1, i, j - 1);
        return false;
    }

    /**
     * Return an inorder traversal of a tree
     */
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> inorder = new ArrayList<Integer>();
        traversalRecurse(root, inorder);
        return inorder;
    }
    public void traversalRecurse(TreeNode root, List<Integer> inorder) {
        if (root != null) {
            traversalRecurse(root, inorder);
            inorder.add(root.val);
            traversalRecurse(root, inorder);
        }
    }

    /**
     * Given n, return the number of unique binary search trees that store 1 - n.
     */
    public int numTrees(int n) {
        int[] mem = new int[n+1];
        // Trees with 0 and 1 nodes each have 1 unique trees
        mem[0] = 1;
        mem[1] = 1;

        for (int i = 2; i <= n; i++) {
            int total = 0;
            for (int j = 1; j <= i; j++) {
                int left = j - 1;
                int right = i - j;
                total += mem[left] * mem[right];
            }
            mem[i] = total;
        }
        return mem[n];
    }

    /**
     * Given a binary tree, determine if it is a valid binary search tree (BST).
     * The left and right subtrees can only contain values that are lesser or greater than root respectively.
     */
    public boolean isValidBST(TreeNode root) {
        return validBSTRecurse(root, null, null);
    }
    public boolean validBSTRecurse(TreeNode root, Integer min, Integer max) {
        if (root == null) {
            return true;
        }
        int val = root.val;
        if (min != null && val <= min) return false;
        if (max != null && val >= max) return false;
        // Make sure all nodes on the left subtree are less than root
        if (!validBSTRecurse(root.left, min, val)) {
            return false;
        }
        // Make sure all the nodes on the right subtree are greater than root
        if (!validBSTRecurse(root.right, val, max)) {
            return false;
        }
        return true;
    }

    /**
     * Do a level order traversal (breadth first search) on root
     * Every level should be its own list appended to this list of lists
     */
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> levels = new ArrayList<>();
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()) {
            List<Integer> level = new ArrayList<>();
            // Recognize that the size of the queue is the size of the level
            int levelSize = queue.size();

            while (levelSize > 0) {
                TreeNode current = queue.poll();
                if (current.left != null) {
                    queue.add(current.left);
                }
                if (current.right != null) {
                    queue.add(current.right);
                }
                level.add(current.val);
            }
            levels.add(level);
        }
        return levels;
    }

    public int maxDepth(TreeNode root) {
        return maxDepthRecurse(root, 1);
    }
    public int maxDepthRecurse(TreeNode root, int depth) {
        if (root == null) {
            return depth - 1;
        }
        return Math.max(maxDepthRecurse(root.left, depth + 1),
                maxDepthRecurse(root.right, depth + 1));
    }

    /**
     * Given preorder and inorder traversal of a tree, construct the binary tree.
     */
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        return null;
    }

    /**
     * Given an m x n matrix, return all elements of the matrix in spiral order.
     */
    public List<Integer> spiralOrder(int[][] matrix) {
        return null;
    }

    /**
     * int[] prices represents stock prices for the ith day
     * Return the maximum amount of profit you can gain
     * You have to buy one then sell one
     */
    public int maxProfit(int[] prices) {
        int minPrice = Integer.MAX_VALUE;
        int maxProfit = 0;
        for (int i = 0; i < prices.length; i++) {
            if (minPrice > prices[i]) {
                minPrice = prices[i];
            }
            else if (prices[i] - minPrice > maxProfit) {
                maxProfit = prices[i] - minPrice;
            }
        }
        return maxProfit;
    }

    public static void main(String[] args) {
        int[] data = {1,2,3};
        System.out.println(subsets(data));
    }
}
