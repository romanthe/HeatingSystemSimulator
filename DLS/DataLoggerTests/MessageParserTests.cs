using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using NUnit.Framework;
using NUnit;
using DataLogger;



namespace DataLoggerTests
{
    [TestFixture]
    public class MessageParserTests
    {
		MessageParser mp;

		[SetUp]
		public void SetUp(){
			mp = new MessageParser("");
		}

        [Test]
        public void GetString_EmptyMessage()
        {
			MessageParser mp = new MessageParser("");

			Assert.Throws(typeof(NoSuchFieldException), () => mp.GetString("gg") );
        }

		[Test]
		public void GetString_MessageWithSingleItem()
		{
			MessageParser mp = new MessageParser("port=1234");

			Assert.AreEqual("1234", mp.GetString ("port"));
		}

		[Test]
		public void GetString_LastField()
		{
			mp.Message = "id=123 port=1236";

			Assert.AreEqual("1236", mp.GetString("port"));
		}

		[Test]
		public void GetString_FirstField(){
			mp.Message = "port=1236 id=545";

			Assert.AreEqual("1236", mp.GetString("port"));
		}

		[Test]
		public void GetInt_ValidMessage(){
			mp.Message = "port=12";

			Assert.AreEqual(12, mp.GetInt("port"));
		}

		[Test]
		public void GetString_NoSuchField(){
			mp.Message = "pp=43";

			Assert.Throws(typeof(NoSuchFieldException), () => mp.GetInt("gg") );
		}

		[Test]
		public void GetInt_NotAnIntException(){
			mp.Message = "port=ab";

			Assert.Throws(typeof(FormatException), () => mp.GetInt("port"));
		}
    }
}
