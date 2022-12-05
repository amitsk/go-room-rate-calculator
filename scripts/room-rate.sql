CREATE TABLE room_rates (
  id              SERIAL PRIMARY KEY,
  zipcode         char(9) unique not null ,
  price 		  numeric(10,2) not null CONSTRAINT positive_price CHECK (price > 0)
);


insert  into room_rates(zipcode, price) values ('97006', 120.50);
insert  into room_rates(zipcode, price) values ('97229', 120.50);
insert  into room_rates(zipcode, price) values ('97103', 150.50);
insert  into room_rates(zipcode, price) values ('97411', 180.50);
insert  into room_rates(zipcode, price) values ('97415', 160.50);
insert  into room_rates(zipcode, price) values ('97110', 170.50);
insert  into room_rates(zipcode, price) values ('97420', 125.50);
insert  into room_rates(zipcode, price) values ('97341', 130.50);
insert  into room_rates(zipcode, price) values ('97439', 140.50);