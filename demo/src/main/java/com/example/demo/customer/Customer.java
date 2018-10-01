@Data
@Entity
public class Customer {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @NotNull
    @Size(min = 2, max = 50, message = "Please provide first size between 2 - 100")
    private String firstName;

    @NotNull
    @Size(min = 2, max = 100, message = "Please provide lastName size between 2 - 100")
    private String lastName;

    @NotNull
    @Min(value = 18, message = "Please provide age >18")
    private Integer age;

    @Email(message = "Please provide valid email address")
    private String email;
}