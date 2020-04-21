using System;
using System.Linq;
using System.Text;

namespace PP0506A
{
    class Program
    {
        static void Main()
        {
            short testsCount = Convert.ToInt16(Console.ReadLine());
            StringBuilder outputBuilder = new StringBuilder();

            for (short c = 0; c < testsCount; c++)
            {
                string line = Console.ReadLine();
                if (line == "")
                {
                    line = Console.ReadLine();
                }

                short pointsCount = Convert.ToInt16(line);
                Tuple<string, int, int, double>[] points = new Tuple<string, int, int, double>[pointsCount];

                for (short i = 0; i < pointsCount; i++)
                {
                    string pointString = Console.ReadLine();
                    string pointName = pointString.Substring(0, pointString.IndexOf(' ') + 1);
                    short[] pointCoordinates = pointString.Remove(0, pointString.IndexOf(' ') + 1).Split(' ').Select(Int16.Parse).ToArray();
                    short x = pointCoordinates[0];
                    short y = pointCoordinates[1];
                    double distanceToOrigin;
                    if (Math.Abs(x) == Math.Abs(y))
                    {
                        distanceToOrigin = Math.Abs(x);
                    }
                    else if (x == 0)
                    {
                        distanceToOrigin = Math.Abs(y);
                    }
                    else if (y == 0)
                    {
                        distanceToOrigin = Math.Abs(x);
                    }
                    else
                    {
                        distanceToOrigin = Math.Sqrt(x * x + y * y);
                    }

                    Tuple<string, int, int, double> point = new Tuple<string, int, int, double>(pointName, x, y, distanceToOrigin);
                    points[i] = point;
                }

                Array.Sort(points, delegate (Tuple<string, int, int, double> point1, Tuple<string, int, int, double> point2) {
                    return point1.Item4.CompareTo(point2.Item4);
                });

                for (short i = 0; i < pointsCount; i++)
                {
                    outputBuilder.AppendFormat("{0} {1} {2}", points[i].Item1, points[i].Item2, points[i].Item3);
                    outputBuilder.AppendLine();
                }

                outputBuilder.AppendLine();
            }

            Console.WriteLine(outputBuilder.ToString());
        }
    }
}
