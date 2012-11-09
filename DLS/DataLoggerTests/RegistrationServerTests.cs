using System;
using NUnit.Framework;
using DataLogger;
using System.Net.Sockets;
using NSubstitute;
using System.Text;
using System.Net;


namespace DataLoggerTests
{
	[TestFixture]
	public class RegistrationServerTests
	{
		IConnection commMock;

		[SetUp]
		public void SetUp(){
			commMock = Substitute.For<IConnection>();
		}

		[TearDown]
		public void TearDown(){
			commMock.ClearReceivedCalls();
		}

		[TestCase("action=register id=name","status=registration_failure")]
		[TestCase("action=register id=name port=1234","status=registered")]
		[TestCase("action=register id=name port=999999", "status=registration_failure")]
		[TestCase("action=register port=1234", "status=registration_failure")]
		public void Process_NoPortRightAnswer (string receiveMessage, string answerMessage){
			RegistrationServer registrationServer = new RegistrationServer(new DataCollector());
			commMock.Read().Returns(receiveMessage);
			commMock.RemoteIp.Returns(IPAddress.Parse("200.0.0.1"));

			registrationServer.Process(commMock);

			commMock.Received().Write(answerMessage);
		}

		[Test]
		public void Process_AddsRightArtifact(){
			IDataCollector dataCollectorMoc = Substitute.For<IDataCollector>();
			RegistrationServer registrationServer = new RegistrationServer(dataCollectorMoc);

			commMock.Read().Returns("action=register id=name port=1234");
			commMock.RemoteIp.Returns(IPAddress.Parse("127.0.0.3"));

			registrationServer.Process(commMock);

			dataCollectorMoc.Received().AddArtifact("name", IPAddress.Parse("127.0.0.3") ,1234);
		}
	}
}

