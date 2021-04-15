package bsep.bsep.security;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;

import org.springframework.core.Ordered;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;

@Component
@Order(Ordered.HIGHEST_PRECEDENCE)
public class XSSFilter implements Filter {

	// Takes user input (HTTP request) and cleans it (removes anything that might be
	// malicious)
	@Override
	public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
			throws IOException, ServletException {
		System.out.println("prva linija");
		XSSRequestWrapper wrappedRequest = new XSSRequestWrapper((HttpServletRequest) request);
		System.out.println("druga linija");
		chain.doFilter(wrappedRequest, response);
		System.out.println("treca linija");
	}
}