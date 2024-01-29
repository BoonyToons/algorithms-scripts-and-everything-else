
using System.Configuration.Assemblies;

class SeletionSort {
    static void Main(string[] args) {
        SelectionSort([2,1,7,1,32,6,3,20]);
        SelectionSort([2,5,2,5,5,3,1,1,2,1]);
        SelectionSort([5,4,3,2,1]);
        SelectionSort([7,3,1,3,2,1,4]);
        SelectionSort([1,2,3,6,5,4]);
        SelectionSort([-6,-5,1,1,2,-1,-1,3,5,2,7]);
    }


    static void SelectionSort(int[] numbers) {

        for(int ii=0; ii<numbers.Length - 1; ii++) {
            int min = numbers[ii];
            int index = ii;
            int firstSlot = numbers[ii];
            for(int kk=ii+1; kk<numbers.Length; kk++) {
                if(min > numbers[kk]) {
                    min = numbers[kk];
                    index = kk;
                }
            }
            numbers[ii] = min;
            numbers[index] = firstSlot; 
        }

        for(int ii=0; ii<numbers.Length; ii++) {
            Console.Write(numbers[ii]);
            Console.Write(" ");
        }
        Console.WriteLine("");
    }
}