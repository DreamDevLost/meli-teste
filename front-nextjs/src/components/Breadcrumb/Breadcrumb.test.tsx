import { expect, test } from 'vitest';
import { render, screen } from '@testing-library/react';
import Breadcrumb from './Breadcrumb';

test('renders breadcrumb links', () => {
    const links = [
        { name: 'Home', url: '/' },
        { name: 'Products', url: '/products' },
        { name: 'Category', url: '/products/category' },
        { name: 'Product', url: '/products/category/product' },
    ];

    render(<Breadcrumb links={links} />);

    const breadcrumbLinks = screen.getAllByRole('link');
    expect(breadcrumbLinks).toHaveLength(links.length);

    breadcrumbLinks.forEach((link, index) => {
        expect(link.textContent).contain(links[index].name);
        expect(link.getAttribute("href")).to.equal(links[index].url);
    });
});