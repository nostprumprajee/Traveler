@Data
public class Customer {
    @Id
    private String id;
    @NotNull
    private String firstName;
    @NotNull
    private String lastName;
    @NotNull
    private Integer age;
    @Email
    private String email;

}