# PubSubInGolang
This project explains how we can send a message packet from producer to Rabbit Queue and read that message packet from Queue to save in DB through consumer.

Steps To Execute The Project

After installing all dependent libraries we need to follow following steps:

1. cd GoAssignment/producer
2. go run publisher.go // This will publish the message packet into HOTEL_INFO queue
3. cd GoAssignment/consumer
4. go run consumer.go // This will read the message packet from HOTEL_INFO queue and write the information into Sqlite database according to different condition.
