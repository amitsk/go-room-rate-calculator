CREATE TABLE IF NOT EXISTS  public.room_rates (
  id              SERIAL PRIMARY KEY,
  zipcode         char(9) unique not null ,
  price 		  numeric(10,2) not null CONSTRAINT positive_price CHECK (price > 0)
);

