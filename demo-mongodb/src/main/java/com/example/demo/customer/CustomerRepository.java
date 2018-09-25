public interface CustomerRepository extends MongoRepository<Customer, String> {
    List<Customer> findByFirstName(String firstName);
    
}