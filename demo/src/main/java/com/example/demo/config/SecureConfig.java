@EnableWebSecurity
@EnableAutoConfiguration
public class SecureConfig extends WebSecurityConfigurerAdapter {
    @Override
    protected void configure(HttpSecurity http) throws Exception {
        http.authorizeRequests().anyRequest().authenticated()
                .and().csrf()
                .disable()
                .httpBasic();
    }
    
}