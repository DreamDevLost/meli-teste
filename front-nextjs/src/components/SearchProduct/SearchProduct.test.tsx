import { render, screen } from "@testing-library/react";
import SearchProduct from "./SearchProduct";
import { describe, expect, test } from "vitest";

describe("SearchProduct", () => {
    const products = [
        {
            id: "1",
            title: "Product 1",
            price: {
                currency: "USD",
                amount: 10,
                decimals: 2,
            },
            picture: "product1.jpg",
        },
        {
            id: "2",
            title: "Product 2",
            price: {
                currency: "USD",
                amount: 20,
                decimals: 2,
            },
            picture: "product2.jpg",
        },
    ];
    render(<SearchProduct product={products} />);

    test("renders product list", () => {
        const productList = screen.getByTestId("product-list");
        expect(productList).toBeInTheDocument();
    });

    test("renders product items", () => {
        const productItems = screen.getAllByTestId("product-item");
        expect(productItems.length).toBe(products.length);
    });

    test("renders product image", () => {
        const productImages = screen.getAllByAltText("Imagem do produto");
        expect(productImages.length).toBe(products.length);
    });

    test("renders product price", () => {
        const productPrices = screen.getAllByText((content, element) =>
            content.startsWith("$")
        );
        expect(productPrices.length).toBe(products.length);
    });

    test("renders product title", () => {
        const productTitles = screen.getAllByTestId("product-title");
        expect(productTitles.length).toBe(products.length);
    });
});