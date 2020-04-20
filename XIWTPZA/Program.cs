using System;
using System.Linq;

namespace XIWTPZA
{
    class Program
    {
        static void Main(string[] args)
        {
            short testsCount = Convert.ToInt16(Console.ReadLine());

            while (testsCount-- > 0)
            {
                string line = Console.ReadLine();
                int[] array = line.Split(' ').Select(Int32.Parse).ToArray();

                long a1;
                long b1;
                if (array[0] > array[1])
                {
                    a1 = array[0];
                    b1 = array[1];
                }
                else
                {
                    a1 = array[1];
                    b1 = array[0];
                }
                double d1 = Math.Sqrt(a1 * a1 + b1 * b1);

                long a2;
                long b2;
                if (array[2] > array[3])
                {
                    a2 = array[2];
                    b2 = array[3];
                }
                else
                {
                    a2 = array[3];
                    b2 = array[2];
                }
                double d2 = Math.Sqrt(a2 * a2 + b2 * b2);



                if (a1 > a2 && b1 > b2)
                {
                    Console.WriteLine("TAK");
                    continue;
                }
                else if (a1 < a2 && b1 < b2)
                {
                    Console.WriteLine("NIE");
                    continue;
                }
                else if (d1 <= d2)
                {
                    Console.WriteLine("NIE");
                    continue;
                }
                else
                {
                    double x1 = Math.Sqrt(Math.Pow(d2 / 2, 2) - Math.Pow(a1 / 2, 2));
                    double y1 = a1 / 2;

                    double x2 = b1 / 2;
                    double y2 = Math.Sqrt(Math.Pow(d2 / 2, 2) - Math.Pow(b1 / 2, 2));

                    double xy = Math.Sqrt(Math.Pow(x1 - x2, 2) + Math.Pow(y1 - y2, 2));

                    if (xy > b2)
                    {
                        Console.WriteLine("TAK");
                        continue;
                    }
                    else
                    {
                        Console.WriteLine("NIE");
                        continue;
                    }
                }
            }
        }
    }
}
