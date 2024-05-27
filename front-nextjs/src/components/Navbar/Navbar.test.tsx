import { render, screen } from "@testing-library/react";
import Navbar from "./Navbar";
import { describe, test, expect } from "vitest";

describe("Navbar", () => {
    render(<Navbar />);

    test("renders logo", () => {
        const logo = screen.getByAltText("Logo");
        expect(logo).toBeInTheDocument();
    });

    test("renders search input", () => {
        const searchInput = screen.getByPlaceholderText("Nunca dejes de buscar");
        expect(searchInput).toBeInTheDocument();
    });

    test("renders search button", () => {
        const searchButton = screen.getByRole("button", { name: "Buscar produtos" });
        expect(searchButton).toBeInTheDocument();
    });
});