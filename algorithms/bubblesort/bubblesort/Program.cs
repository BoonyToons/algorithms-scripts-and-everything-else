
using System.Globalization;

class Program {
    
    static void Main(string[] args) {
        
        BubbleSort([5,4,3,2,1]);
        BubbleSort([2,65,2,1,8,4,3,1]);
        BubbleSort([2,3,1,5,6,4]);
        BubbleSort([-1,3,-6,2,1,-3,5]);
        BubbleSort([]);
    }

    static void BubbleSort(int[] numbers){
        for(int ii=0; ii<numbers.Length; ii++) {
            for(int kk=0; kk<numbers.Length - 1; kk++){
                if(numbers[kk] > numbers[kk+1]) {
                    int left = numbers[kk];
                    int right = numbers[kk+1];
                    numbers[kk] = right;
                    numbers[kk+1] = left;
                }
            }
        }

        for(int ii=0; ii<numbers.Length; ii++) {
            Console.Write(numbers[ii]);
            Console.Write(" ");
        }
        Console.WriteLine("");
    }
}
