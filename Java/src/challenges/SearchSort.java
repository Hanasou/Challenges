package challenges;

import java.util.Arrays;

public class SearchSort {

    public static int binarySearchIterative(int[] arr, int target) {
        int start = 0;
        int end = arr.length - 1;
        while (start <= end) {
            int mid = (start + end) / 2;
            int val = arr[mid];
            if (target == val) {
                return mid;
            }
            if (target > val) {
                // We want to assign start to mid + 1 because we know that
                // target is after mid
                start = mid + 1;
            }
            else {
                // We want to assign mid to mid - 1 because we know that
                // target is before mid
                end = mid - 1;
            }
        }
        return -1;
    }
    public static int binarySearchRecursive(int[] arr, int target) {
        return searchRecurse(arr, target, 0, arr.length - 1);
    }
    public static int searchRecurse(int[] arr, int left, int right, int target) {
        if (left <= right) {
            int mid = (right + left) / 2;
            int val = arr[mid];
            if (target > val) {
                return searchRecurse(arr, mid + 1, right, target);
            } else {
                return searchRecurse(arr, left, mid - 1, target);
            }
        }
        return -1;
    }

    public void mergesort(int[] arr) {
        int[] helper = new int[arr.length];
        mergesortRecurse(arr, helper, 0, arr.length - 1);
    }
    public void mergesortRecurse(int[] arr, int[] helper, int left, int right) {
        // We're done once we have a single element array
        if (left < right) {
            int middle = (left + right) / 2;
            // sort the left half
            mergesortRecurse(arr, helper, left, middle);
            // sort the right half
            mergesortRecurse(arr, helper, middle + 1, right);
            // merge
            merge(arr, helper, left, middle, right);
        }
    }
    // This function should merge two arrays together
    // The ints left and right denote which arrays we're sorting
    public void merge(int[] arr, int[] helper, int left, int mid, int right) {
        // We're going to allocate some extra space for a helper array
        // To keep values for the merging process
        for (int i = left; i <= right; i++) {
            helper[i] = arr[i];
        }

        // Initialize helper pointers
        int hleft = left;
        int hright= right;
        // current denotes the index of the spot we're currently modifying
        int current = left;

        // start merging
        // We need a helper array because we're replacing elements in our original array
        while(hleft <= mid && hright <= right) {
            if (helper[hleft] <= helper[hright]) {
                arr[current] = helper[hleft];
                hleft++;
            } else {
                arr[current] = helper[hright];
                hright++;
            }
            current++;
        }

        // Fill in values for the array that still contains them
        // In this particular example, we're only filling in values from the left half because
        // values from the right half are already there.
        int remaining = mid - hleft;
        for (int i = 0; i <= remaining; i++) {
            arr[current + i] = helper[hleft + i];
        }
    }

    public static void quicksort(int[] arr) {
        quicksortHelp(arr, 0, arr.length - 1);
    }
    public static void quicksortHelp(int[] arr, int left, int right) {
        int index = partition(arr, left, right);
        // I guess this is only a thing if we're selecting the middle value as the partition index?
        // Otherwise I don't really know why this exists.
        if (left < index - 1) {
            quicksortHelp(arr, left, index - 1);
        }
        if (index < right) {
            quicksortHelp(arr, index, right);
        }
    }
    public static int partition(int[] arr, int start, int end) {
        int pivot = arr[(start + end) / 2];
        while (start <= end) {
            while (arr[start] < pivot) {
                start++;
            }
            while (arr[end] > pivot) {
                end--;
            }
            // Start could be after end after those two loops above are done
            // That means everything is already in place
            if (start <= end) {
                int temp = arr[start];
                arr[start] = arr[end];
                arr[end] = temp;
                start++;
                end--;
            }
        }
        System.out.println(start);
        return start;
    }

    public static void main(String[] args) {
        int[] data = {1,2,3,4,5};
        quicksort(data);
        System.out.println(Arrays.toString(data));
    }
}
