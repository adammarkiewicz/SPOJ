using System;
using System.Collections.Generic;
using System.Text;

namespace PTCLTZ
{
    class Program
    {
        static void Main()
        {
            int testsCount = Convert.ToInt16(Console.ReadLine());
            StringBuilder output = new StringBuilder();

            while (testsCount-- > 0)
            {
                int s = Convert.ToInt16(Console.ReadLine());
                Dictionary<int, int> map = new Dictionary<int, int>();
                int counter = 0;

                while (s != 1)
                {
                    if (map.TryGetValue(s, out int value))
                    {
                        s = value;
                    }
                    else
                    {
                        value = Collatz(s);
                        map.Add(s, value);
                        s = value;
                    }
                    counter++;
                }
                output.Append(counter + Environment.NewLine);
            }

            Console.WriteLine(output.ToString());
        }

        static int Collatz(int x)
        {
            if (x % 2 == 0)
            {
                return x / 2;
            }
            else
            {
                return (3 * x) + 1;
            }
        }
    }
}
