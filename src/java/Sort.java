package java;

/**
 * 七种排序算法Java版
 * @author Huangxiaoxu
 *
 */
public class Sort {

    /**
     * 打印数组
     *
     * @param data
     */
    public static void displayData(int[] data) {
        for (int d : data) {
            System.out.print(d + " ");
        }
        System.out.println();
    }

    /**
     * 冒泡排序算法，时间复杂度O(n2)，算法具有稳定性，堆排序和快速排序算法不具有稳定性，即排序后相同元素的顺序会发生变化
     *
     * @param src
     */
    public static void bubbleSort(int[] src) {
        if (src.length > 0) {
            int length = src.length;
            for (int i = 1; i < length; i++) {
                for (int j = 0; j < length - i; j++) {
                    if (src[j] > src[j + 1]) {
                        int temp = src[j];
                        src[j] = src[j + 1];
                        src[j + 1] = temp;
                    }
                }
            }
        }
    }

    /**
     * 改进的冒泡排序算法
     *
     **/
    public static void improvedBubbleSort(int[] src) {
        if (src.length > 0) {
            int length = src.length;
            boolean flag = true;
            for (int i = 1; i < length && flag; i++)
            {
                flag = false;
                for (int j = 0; j < length - i; j++) {
                    if (src[j] > src[j + 1]) {
                        int temp = src[j];
                        src[j] = src[j + 1];
                        src[j + 1] = temp;
                        flag = true;
                    }
                }
            }
        }
    }

    /**
     * 快速排序，时间复杂度O(nlogn)，最坏时间复杂度O(n2)，平均时间复杂度O(nlogn)，算法不具稳定性
     *
     * @param src
     * @param begin
     * @param end
     */
    public static void quickSort(int[] src, int begin, int end) {
        int pos;
        if (begin < end) {
            pos = partition(src,begin,end);
            quickSort(src, begin, pos - 1);
            quickSort(src, pos + 1, end);
        }
    }

    /**
     * 分治思想，查找支点
     * @param data
     * @param low
     * @param high
     * @return
     */
    private static int partition(int data[],int low,int high){
        int key = data[low];
        while(low<high){
            while(low<high && data[high]>key) //从右往左
                high--;
            data[low] = data[high];

            while(low<high && data[low]<key)  //从左往右
                low++;
            data[high] = data[low];
        }
        data[low] = key;
        return low;
    }

    /**
     * 选择排序，分为简单选择排序、树形选择排序(锦标赛排序)、堆排序 此算法为简单选择排序
     *
     * @param a
     */
    public static void selectSort(int[] a) {
        int length = a.length;
        for (int i = 0; i < length; i++) {
            int minIndex = i;
            for (int j = i + 1; j < a.length; j++) {
                if (a[j] < a[minIndex]) {
                    minIndex = j;
                }
            }
            if (minIndex != i) {
                int temp = a[minIndex];
                a[minIndex] = a[i];
                a[i] = temp;
            }
        }
    }

    /**
     * 插入排序，适用于少量数据的排序，时间复杂度O(n2)，是稳定的排序算法，原地排序
     *
     * @param a
     */
    public static void insertSort(int[] a) {
        int length = a.length;
        for (int i = 1; i < length; i++) {
            int temp = a[i];
            int j = i;
            for (; j > 0 && a[j - 1] > temp; j--) {
                a[j] = a[j - 1];
            }
            a[j] = temp;
        }
    }

    /**
     * 归并排序算法，稳定排序，非原地排序，空间复杂度O(n)，时间复杂度O(nlogn)
     *
     * @param a
     * @param low
     * @param high
     */
    public static void mergeSort(int a[], int low, int high) {
        if (low < high) {
            mergeSort(a, low, (low + high) / 2);
            mergeSort(a, (low + high) / 2 + 1, high);
            merge(a, low, (high + low) / 2, high);
        }
    }

    /**
     * 归并排序辅助方法，合并
     *
     * @param a
     * @param low
     * @param mid
     * @param high
     */
    private static void merge(int[] a, int low, int mid, int high) {
        int[] b = new int[high - low + 1];
        int s = low;
        int t = mid + 1;
        int k = 0;
        while (s <= mid && t <= high) {
            if (a[s] <= a[t])
                b[k++] = a[s++];
            else
                b[k++] = a[t++];
        }
        while (s <= mid)
            b[k++] = a[s++];
        while (t <= high)
            b[k++] = a[t++];
        for (int i = 0; i < b.length; i++) {
            a[low + i] = b[i];
        }
    }

    /**
     * 希尔排序的一种实现方法
     *
     * @param a
     */
    public static void shellSort(int[] a) {
        int temp;
        for (int k = a.length / 2; k > 0; k /= 2) {
            for (int i = k; i < a.length; i++) {
                for (int j = i; j >= k; j -= k) {
                    if (a[j - k] > a[j]) {
                        temp = a[j - k];
                        a[j - k] = a[j];
                        a[j] = temp;
                    }
                }
            }
        }
    }

    /**
     * 堆排序，最坏时间复杂度O(nlog2n)，平均性能接近于最坏性能。由于建初始堆所需的比较次数多，故堆不适合记录较少的比较 堆排序为原地不稳定排序
     *
     * @param array
     */
    public static void heapSort(int[] array) {
        for (int i = 1; i < array.length; i++) {
            makeHeap(array, i);
        }

        for (int i = array.length - 1; i > 0; i--) {
            int temp = array[i];
            array[i] = array[0];
            array[0] = temp;
            rebuildHeap(array, i);
        }
    }

    /**
     * 堆排序辅助方法---创建堆
     *
     * @param array
     * @param k
     */
    private static void makeHeap(int[] array, int k) {
        int current = k;
        while (current > 0 && array[current] > array[(current - 1) / 2]) {
            int temp = array[current];
            array[current] = array[(current - 1) / 2];
            array[(current - 1) / 2] = temp;
            current = (current - 1) / 2;
        }
    }

    /**
     * 堆排序辅助方法---堆的根元素已删除，末尾元素已移到根位置，开始重建
     *
     * @param array
     * @param size
     */
    private static void rebuildHeap(int[] array, int size) {
        int currentIndex = 0;
        int right = currentIndex * 2 + 2;
        int left = currentIndex * 2 + 1;
        int maxIndex = currentIndex;
        boolean isHeap = false;
        while (!isHeap) {
            if (left < size && array[currentIndex] < array[left]) {
                maxIndex = left;
            }
            if (right < size && array[maxIndex] < array[right]) {
                maxIndex = right;
            }
            if (currentIndex == maxIndex) {
                isHeap = true;
            } else {
                int temp = array[currentIndex];
                array[currentIndex] = array[maxIndex];
                array[maxIndex] = temp;
                currentIndex = maxIndex;
                right = currentIndex * 2 + 2;
                left = currentIndex * 2 + 1;
            }
        }
    }

    public static void main(String[] args) {
        int data[] = { 2, -1, 5, 4, 6, 8, 7, -3 };
        Sort.displayData(data);
        Sort.bubbleSort(data);
        Sort.displayData(data);
    }
}
