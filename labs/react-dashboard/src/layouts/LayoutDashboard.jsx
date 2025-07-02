import { Box } from "@mui/material";
import { Outlet } from "react-router";

const LayoutDashboard = () => {
    return (
        <Box sx={{
            display: 'flex',
            position: 'relative',
        }}>
            <Outlet />
        </Box>
    );
}

export default LayoutDashboard;