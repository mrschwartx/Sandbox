import { useState, useEffect } from "react";
import { useNavigate } from "react-router";
import { toast } from "react-toastify";
import {
  Box,
  Button,
  CircularProgress,
  IconButton,
  InputAdornment,
  Link,
  Stack,
  TextField,
  useMediaQuery,
} from "@mui/material";
import PersonIcon from "@mui/icons-material/Person";
import EmailIcon from "@mui/icons-material/Email";
import LockIcon from "@mui/icons-material/Lock";
import Visibility from "@mui/icons-material/Visibility";
import VisibilityOff from "@mui/icons-material/VisibilityOff";

import Logo from "../../assets/images/icons/react.svg";
import validateInput from "../../utils/validateInput";
import RegisterApi from "./RegisterApi";

const INITIAL_VALUE = {
  name: "",
  email: "",
  password: "",
  showPassword: false,
};

const INITIAL_ERROR = {
  name: "",
  email: "",
  password: "",
};

const RegisterForm = () => {
  const navigate = useNavigate();

  const tablet = useMediaQuery("(max-width: 920px)");

  const [loading, setLoading] = useState(false);
  const [allowSubmit, setAllowSubmit] = useState(false);

  const [values, setValues] = useState(INITIAL_VALUE);
  const [error, setError] = useState(INITIAL_ERROR);

  const valid = Boolean(error.email === "" && error.password === "");

  const handleChange = (e) => {
    // eslint-disable-next-line no-useless-escape
    const emailPattern = /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g;
    setError(INITIAL_ERROR);

    const { name, value } = e.target;
    setValues({ ...values, [name]: value });

    if (name === "email") {
      if (!emailPattern.test(value)) {
        setError({ ...error, email: "Invalid email!" });
      } else {
        setError({ ...error, email: "" });
      }
    } else {
      const { fieldName, errorMsg } = validateInput(e);
      if (errorMsg) {
        setError({ ...error, [fieldName]: errorMsg });
      } else {
        setError(INITIAL_ERROR);
      }
    }
  };

  const handleClickShowPassword = () => {
    setValues({ ...values, showPassword: !values.showPassword });
  };

  const handleMouseDownPassword = (event) => {
    event.preventDefault();
  };

  useEffect(() => {
    const isFormFilled = Boolean(values.email !== "" && values.password !== "");
    setAllowSubmit(isFormFilled);
  }, [values]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = { ...values };
    delete formData.showPassword;
    const { name, email, password } = formData;

    if (name === "") {
      setError({
        ...error,
        name: "Name is required",
      });
    }

    if (email === "") {
      setError({
        ...error,
        email: "Email is required",
      });
    }

    if (password === "") {
      setError({
        ...error,
        password: "Password is required",
      });
    }

    if (valid && allowSubmit) {
      setLoading(true);

      const response = await RegisterApi(formData);
      
      setLoading(false);

      if (response.success) {
        navigate("/login");
      } else {
        toast.error(response.error || "Registration failed.");
      }
    }
  };

  return (
    <Box
      sx={{
        width: tablet ? "100%" : "60vw",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        px: !tablet && 18,
      }}
    >
      <img
        onClick={() => navigate("/")}
        style={{
          width: 80,
          cursor: "pointer",
          marginBottom: 16,
        }}
        src={Logo}
        alt="logo-vaksin-id"
      />
      <Stack
        spacing={3}
        sx={{ width: tablet ? "75%" : "100%", alignItems: "center" }}
      >
        <TextField
          label="Name"
          name="name"
          placeholder="Your full name"
          error={error.name !== ""}
          helperText={error.name}
          value={values.name}
          onChange={handleChange}
          variant="outlined"
          fullWidth
          required
          autoFocus
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <PersonIcon />
              </InputAdornment>
            ),
          }}
        />
        <TextField
          label="Email"
          name="email"
          placeholder="Your email address"
          error={error.email !== ""}
          helperText={error.email}
          value={values.email}
          onChange={handleChange}
          variant="outlined"
          fullWidth
          required
          autoFocus
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <EmailIcon />
              </InputAdornment>
            ),
          }}
        />
        <TextField
          label="Password"
          name="password"
          placeholder="Your Password"
          error={error.password !== ""}
          helperText={error.password}
          value={values.password}
          onChange={handleChange}
          type={values.showPassword ? "text" : "password"}
          variant="outlined"
          fullWidth
          required
          InputProps={{
            startAdornment: (
              <InputAdornment position="start">
                <LockIcon />
              </InputAdornment>
            ),
            endAdornment: (
              <InputAdornment position="end">
                <IconButton
                  aria-label="toggle password visibility"
                  onClick={handleClickShowPassword}
                  onMouseDown={handleMouseDownPassword}
                >
                  {values.showPassword ? <VisibilityOff /> : <Visibility />}
                </IconButton>
              </InputAdornment>
            ),
          }}
        />
        <Link element="a" href="#" sx={{ alignSelf: "end" }}>
          Forgot your password?
        </Link>
        <Button
          type="submit"
          disabled={!allowSubmit || loading}
          onClick={(e) => {
            handleSubmit(e);
          }}
          variant="contained"
          sx={{ height: 56 }}
          fullWidth
          startIcon={
            loading ? <CircularProgress size={24} thickness={5} /> : undefined
          }
        >
          Register
        </Button>
      </Stack>
    </Box>
  );
};

export default RegisterForm;
