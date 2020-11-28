/*
There are n soldiers standing in a line.
Each soldier is assigned a unique rating value.
You have to form a team of 3 soldiers amongst them under the following rules.
- Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k])
- A team is valid if: (rating[i] < rating[j] < rating[k]) or
(rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n)
Return the number of teams you can form with this condition
*/
function numTeamsBruteForce(rating: number[]): number {
    let validTeams = 0
    for (let i = 0; i < rating.length; i++) {
        for (let j = i + 1; j < rating.length; j++) {
            for (let k = j + 1; k < rating.length; k++) {
                let cond1 = rating[i] < rating[j] && rating[j] < rating[k]
                let cond2 = rating[i] > rating[j] && rating[j] > rating[k]
                if (cond1 || cond2) {
                    console.log(rating[i], rating[j], rating[k])
                    validTeams++
                }
            }
        }
    }
    return validTeams
}

function numTeams(rating: number[]): number {
    // Take kind of a dp approach here.
    // We're going to make arrays that store the number of elements
    // inside rating that are greater than or less than rating[i] respectively
    // for each index.
    // For example, g and s are the arrays that contain greater than and less than values respectively
    // g[i] holds the number of elements that are greater than rating[i] in the rating array
    // s[i] is the same but for values less than rating[i]
    let n = rating.length
    let g: number[] = Array.apply(null, new Array(n)).map(()=> 0);
    let s: number[] = Array.apply(null, new Array(n)).map(()=> 0);
    
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            if (rating[i] < rating[j]) {
                g[i]++
            } else {
                s[i]++
            }
        }
    }
    // Compare all the values in g[i] and s[i] to another index to get the final answer
    // e.g. take array [2,5,3,4,1], g = [3,0,1,0,0], s = [1,3,1,1,0]
    // There are no values greater than 5 after it, so we can't make a team with 2 and 5
    // There is one value greater than 3 after it, so we can make a team of 2, 3, and whatever the other number is
    let answer = 0
    for (let i = 0; i < n; i++) {
        for (let j = i + 1; j < n; j++) {
            if (rating[j] > rating[i]) {
                answer += g[j]
            } else {
                answer += s[j]
            }
        }
    }
    return answer
}

/*
Given a non-empty integer array, find the minimum number of moves required to
make all array elements equal where a move is incrementing a selected value by 1
or decrementing a selected value by 1
*/
function minMoves2(nums: number[]): number {
    // We probably need the mean of the array
    // Then we need to figure out how many increments/decrements
    // it takes to get each element to each that (which is just the difference between
    // the two elements)
    let answer = 0;
    const sorted = nums.sort();
    const mid = Math.floor(sorted.length / 2);
    let median = sorted[mid];
    for (let n of nums) {
        answer += Math.abs(n - median);
    }
    return answer;
}

/*
In the "100 game", two players take turns adding, to a running total, any integer between 1 and 10.
The first player who causes the running total to reach or exceed 100 wins.
What if we change the game so that players cannot reuse integers?
If a player can force a win, then return true
If they can't, then return false.
Example: maxChoosableInteger = 10, desiredTotal = 11
Expected Output: False
The first player cannot win in this situation. Any number they choose, the second player
can choose something else and win
Assume that both players play optimally
*/
function canIWin(maxChoosableInteger: number, desiredTotal: number): boolean {
    /*
    This is very obviously a recursion problem
    Should I be brave and do a DP approach? Fuck that.
    */
   let availableNumbers = [];
   // let's just do this the old fashioned way
   // Imperative programming I guess
   for (let i = 1; i <= maxChoosableInteger; i++) {
       availableNumbers.push(i);
   }
   // Helper function
   function canWinHelper(availableNumbers: number[], desiredTotal: number, player: boolean): boolean {
       // This condition is too long to put in my if statements
       let cond1 = availableNumbers[availableNumbers.length - 1] >= desiredTotal;
       // First player's win condition
       if (cond1 === true && player === true) {
           return true;
       }
       for (let n of availableNumbers) {
           let filtered = availableNumbers.filter(i => i != n);
           if (filtered[filtered.length - 1] > desiredTotal - n && player === false) {
               continue;
           }
           return canWinHelper(filtered, desiredTotal - n, !player);
       }
       return false;
   }
   return canWinHelper(availableNumbers, desiredTotal, true);
}

/**
 * Given an integer array nums, find the sum of the elements between i and j (i <= j) inclusive
 * The update function modifies nums by updating the element at index i to val
 * Constraints:
 * Array is only modifiable with the update function
 * You may assume that the calls to update and sumRange are distributed evenly
 * 0 <= i <= j <= nums.length - 1
 */
class NumArray {
    private nums: number[];
    
    constructor(nums: number[]) {
        this.nums = nums;
    }

    update(i: number, val: number): void {

    }

    sumRange(i: number, j: number): number {
        return 0;
    }
}

console.log(canIWin(4, 6));