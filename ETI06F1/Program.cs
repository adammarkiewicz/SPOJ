using System;
using System.Linq;

namespace ETI06F1
{
    class Program
    {
        static void Main()
        {
            double[] array = Console.ReadLine().Split(' ').Select(Double.Parse).ToArray();

            double r = array[0];
            double d = array[1];

            Console.WriteLine(Math.PI * (r * r - d * d / 4));
        }
    }
}