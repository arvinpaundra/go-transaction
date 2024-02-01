CREATE TABLE IF NOT EXISTS order_details (
  id bigint auto_increment primary key,
  order_id bigint not null,
  product_name varchar(100) not null,
  qty integer not null,
  price float not null,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  constraint foreign key (order_id) references orders(id)
);