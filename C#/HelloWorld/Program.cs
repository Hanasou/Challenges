// 'using' I assume basically means import
// Putting this here means that we don't have to write "System.Console.WriteLine"
using System;
using System.Linq;
using System.Collections.Generic;

namespace C_
{
    // Template for creating custom exceptions in case I want to do that
    // Class : Inherited Class is the syntax in case that's confusing
    public class CustomException : Exception {

        public CustomException(string message, Exception innerException) 
            :base(message, innerException) {
            // I suppose base() is basically super()
            
        }
    }
    // Class declaration nonsense
    public class Human {
        // It appears that getters and setters can be defined on the same line
        // Having these set here determines whether you can read/write it
        // Don't have to modify these variables with methods, just access/change them directly
        // object.Name to get it
        // object.Name = "New Name" to set it
        public string Name {get; set;}
        public int Age {get; set;}

        // Constructor works like I would expect it to
        public Human(string name, int age) {
            this.Name = name;
            this.Age = age;
        }

        public override string ToString() {
            return this.Name + " " + this.Age;
        }
    }

    // C# has structs too
    // Structs are value types, Classes are reference types.
    // That means that structs contain values while objects are pointers
    public struct Coords {
        public double X {get; set;}
        public double Y {get; set;}

        public Coords(double x, double y) {
            this.X = x;
            this.Y = y;
        }
    }
    class Program
    {
        // delegates are kind of like method fields in an interface
        // Use them to pass methods are arguments to other methods
        // You can assign functions that match the syntax defined by the delegate
        // to an instance of the delegate and use the delegate to call the function.
        public delegate int NumberHandler(int n);

        public static int AddOne(int n) {
            return n + 1;
        }

        public static int MultThree(int n) {
            return n * 3;
        }

        public static void TestDelegates() {
            // Pass in my function into the delegate I defined
            NumberHandler handler = AddOne;
            // Does this only call the last one?
            handler += MultThree;
            Console.WriteLine(handler(1));
        }

        static void LinqTest() {
            var listOfPeople = new List<Human>();
            var peopleNames = new List<String> {"These", "Are", "Not", "Actually", "Names"};
            var peopleAges = new List<int> {20, 32, 12, 25, 17};
            for(int i = 0; i < peopleNames.Count(); i++) {
                listOfPeople.Add(new Human(peopleNames[i], peopleAges[i]));
            }
            
            // Where() is basically filter
            var underTwenty = listOfPeople.Where(h => h.Age < 20);
            foreach (var person in underTwenty) {
                Console.WriteLine(person.ToString());
            }
            Console.WriteLine(" ");

            // We can sort by particular fields with OrderBy()
            var sortByAge = listOfPeople.OrderBy(h => h.Age);
            // Also if you haven't figured it out already, we can chain operations
            // Get all the humans in the list and add their age by 5
            // Select is used for projection, if you just want to apply transformations
            // just use ForEach()
            var mapAndOrderBy = listOfPeople
                                            .OrderBy(h => h.Age)
                                            .Select(h => h.Age);
            
            // There's another syntax called LINQ query operators
            // This is more like SQL if you want to do that
            foreach (var person in mapAndOrderBy) {
                Console.WriteLine(person.ToString());
            }
            
            // You can grab a single item with Single()
            // If you're not sure if it exists or not, use SingleOrDefault()
            var singlePerson = listOfPeople.SingleOrDefault(h => h.Name == "Actually");
            // First() gives you the first item
            var firstPerson = listOfPeople.First(h => h.Name == "Actually");
            // Last() and LastOrDefault() also exist. Hopefully I know what they do at this point.
            
            // Skip() and Take() allows you to paginage your queries
            // This query skips the first two and takes the next three.
            var skipAndTake = listOfPeople.Skip(2).Take(3);
            // Aggregation methods also exist like Count(), Max(), Min(), Sum(), Average(), etc
        }

        // Lambda functions exist
        static void LambdaTest() {
            // args => expression
            // lambda expression for squaring a number
            // n => n * n;
            // We need to point to our lambda expressions with a delegate
            // Func<int, int> delegate is used for functions with a return type
            // (<int, int>)the first is the type that we're taking in, the second is the return type
            // So this would be the full lambda expression declaration
            Func<int, int> square = n => n*n;
            Console.WriteLine(square(5));
        }

        static void NullableNonsense() {
            // Reference types can be null (objects)
            // Value types cannot be null (structs and primitives)
            // In order to have a value type be null
            // you need to wrap it with a Nullable type
            Nullable<int> nullInt = null;
            // There's a shorthand way to do this too
            int? otherNullInt = null;
            // Nullable types have fields and methods associated with them
            // GetValueOrDefault(), HasValue, Value
            // Use the first two if you're unsure if the Nulltable type is null or not
            // If you call Value on a null value then you'll get an exception
            Console.WriteLine(nullInt);
            Console.WriteLine(otherNullInt);
        }

        public static void PlayingWithStructs() {
            // Instantiated just like objects
            Coords p = new Coords(32, 42);
        }

        public static void ListNonsense() {
            // List declaration
            var names = new List<String> {"These", "Are", "Not", "Names"};
            var empty = new List<int>();
            // Foreach loop
            foreach (var name in names) {
                Console.WriteLine(name);
            }
            // You can call add to add stuff to list
            names.Add("Also not a name");
            // Sort and IndexOf does what you expect them to do.
            // Sort() sorts in place like always

            // Just make sure that empty list declaration is correct
            for (int i = 0; i < 12; i++) {
                empty.Add(i);
            }
            // the Sum() function is part of System.Linq
            Console.WriteLine(empty.Sum());
        }

        public static void PlayingWithPeople() {
            // Instantiating an object works like I would expect it to
            Human man = new Human("Roy", 32);
            Console.WriteLine(man.Name + " " + man.Age);
        }

        public static void PlayingWithMaps() {
            // Declaration works how I would expect it to
            var dict = new Dictionary<string, int>();
            // Indexing and getting items works like how it does in modern programming
            // languagues. Methods like get and put are also supported I think but
            // why would I use them.
            dict["One"] = 1;
            dict["Two"] = 2;
            dict["Three"] = 3;
            Console.WriteLine(dict["One"]);
            Console.WriteLine(dict["Two"]);
            // Some important things to note
            // - Dictionaries do not allow null keys
            // - Dictionary will replace existing value with new one with [] indexing
            // but will throw ArgumentException if you use Add()
            // - Dictionary will throw KeyNotFoundException if you try to get a non-existant key
            // you can use TryGetValue method instead of [] indexing to avoid this
        }

        // Here's a function with a constraint
        // We're taking in two generic types, so the compiler treats them as objects
        // We don't know if the types can be compared to each other or not, so
        // we have to add a constraint that both of these types have to implement the IComparable
        // interface. That's what "where T : IComparable" means.
        public static T Max<T>(T a, T b) where T : IComparable {
            // Shorthand if/else statement that's featured in other languages as well
            // basically the equivalent of
            // if (a.CompareTo(b) > 0) return a else return b
            return a.CompareTo(b) > 0 ? a : b;
        }

        // Default is a pretty useful keyword
        public static T GetDefault<T>() {
            return default(T);
        }

        static void Main(string[] args)
        {
            LinqTest();
        }
    }
}
