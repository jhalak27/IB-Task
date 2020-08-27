import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Navbar,Nav } from 'react-bootstrap'

export default function NavBar() {
    return(
        <Navbar bg="dark" variant="dark">
        <Navbar.Brand href="/">Interview Scheduling</Navbar.Brand>
        <Nav className="mr-auto">
        <Nav.Link href="/">Create</Nav.Link>
        <Nav.Link href="/meeting">Edit</Nav.Link>
        {/* <Nav.Link href="#pricing">Pricing</Nav.Link> */}
        </Nav>
        {/* <Form inline>
        <FormControl type="text" placeholder="Search" className="mr-sm-2" />
        <Button variant="outline-info">Search</Button>
        </Form> */}
    </Navbar>
    );
}