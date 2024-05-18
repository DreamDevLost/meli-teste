import Navbar from "@components/Navbar/Navbar.tsx";
import './App.scss'
import Breadcrumb from "@components/Breadcrumb/Breadcrumb.tsx";
import SearchProduct from "@components/SearchProduct/SearchProduct.tsx";

function App() {
    return (
        <>
            <Navbar/>
            <main>
                <Breadcrumb links={[
                    {name: 'qwr', url: 'http://localhost:8080'},
                    {name: 'qwr', url: 'http://localhost:8080'},
                    {name: 'qwr'},
                ]}/>
                <SearchProduct/>
            </main>
        </>
    )
}

export default App
