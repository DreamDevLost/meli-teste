import { render } from '@testing-library/react';
import Product from './Product';
import { describe, it, expect } from 'vitest';

describe('Product component', () => {
    const item = {
        id: '1',
        title: 'Example Product',
        price: {
            currency: 'USD',
            amount: 9.99,
            decimals: 2,
        },
        picture: 'example.jpg',
        condition: 'New',
        free_shipping: true,
        sold_quantity: 10,
        description: 'This is an example product.',
    };

    it('renders product details correctly', () => {
        const { getByAltText, getByText } = render(<Product item={item} />);

        expect(getByAltText('Product')).toBeInTheDocument();
        expect(getByText('Example Product')).toBeInTheDocument();
        expect(getByText('$ 9.99')).toBeInTheDocument();
        expect(getByText('Comprar')).toBeInTheDocument();
        expect(getByText('Descrição do produto')).toBeInTheDocument();
        expect(getByText('This is an example product.')).toBeInTheDocument();
    });
});