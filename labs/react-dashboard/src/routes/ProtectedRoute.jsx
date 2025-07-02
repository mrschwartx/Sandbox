import { useEffect } from "react";
import { useState } from "react";
import { Navigate, Outlet, useLocation } from "react-router";
import { Helmet } from "react-helmet";

import AuthService from "../services/AuthService";
import { getHeading } from "../utils/getHeading";

const TITLE = "Dashboard";

// User cannot access, if user has been login
const ProtectedRoute = () => {
  const [title, setTitle] = useState(TITLE);
  const { pathname } = useLocation();

  useEffect(() => {
    const heading = getHeading(pathname)
    setTitle(heading || TITLE);
  }, [pathname]);

  if (AuthService.isAuthorized()) {
    return <Navigate to="/" replace />;
  }

  return (
    <>
      <Helmet>
        <meta charSet="utf-8" />
        <title>{ title }</title>
      </Helmet>
      <Outlet />
    </>
  );
};

export default ProtectedRoute;
