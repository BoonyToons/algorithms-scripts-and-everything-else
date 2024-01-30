
using System.Configuration.Assemblies;

class LinearSearch {
    static void Main(string[] args) {
        Linear([1,5,2,7,8,2], 5);
        Linear([1,5,2,7,8,2], 8);
        Linear([1,5,2,7,8,2], 1);
        Linear([1,5,2,7,8,2], 0);
    }

    static void Linear(int[] numbers, int num) {
        for(int ii=0; ii<numbers.Length; ii++) {
            if(num == numbers[ii]) {
                Console.WriteLine(num);
                break;
            }
        }
    }
}