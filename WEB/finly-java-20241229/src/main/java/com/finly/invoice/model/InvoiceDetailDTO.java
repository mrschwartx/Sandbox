package com.finly.invoice.model;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.finly.customer.model.CustomerDTO;
import com.finly.user.model.UserDTO;
import jakarta.validation.constraints.Digits;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Size;
import lombok.Getter;
import lombok.Setter;

import java.math.BigDecimal;
import java.util.Date;

@Getter
@Setter
public class InvoiceDetailDTO {

    private Long id;

    @NotNull(message = "Amount cannot be null.")
    @Digits(integer = 10, fraction = 2, message = "Amount must be a valid number with up to 10 digits and 2 decimal places.")
    @JsonFormat(shape = JsonFormat.Shape.STRING)
    private BigDecimal amount;

    @NotNull(message = "Status cannot be null.")
    @Size(max = 255, message = "Status cannot exceed 255 characters.")
    private String status;

    @NotNull(message = "Due date cannot be null.")
    @JsonFormat(pattern = "MM/dd/yyyy")
    private Date dueDate;

    @NotNull(message = "Owner information is required.")
    private UserDTO owner;

    @NotNull(message = "Customer information is required.")
    private CustomerDTO customer;
}