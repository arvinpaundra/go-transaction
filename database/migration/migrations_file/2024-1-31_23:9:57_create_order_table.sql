CREATE TABLE IF NOT EXISTS orders (
  id bigint auto_increment primary key,
  user_id bigint not null,
  shipment varchar(100) not null,
  status enum('created', 'packed', 'picked', 'shipped', 'delivered') default 'created',
  grand_total float not null,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp,
  constraint foreign key (user_id) references users(id)
);