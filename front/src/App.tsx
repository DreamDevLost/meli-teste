import Navbar from "@components/Navbar/Navbar.tsx";
import './App.scss'
import Breadcrumb from "@components/Breadcrumb/Breadcrumb.tsx";

function App() {
    return (
        <>
            <Navbar/>
            <main>
                <Breadcrumb links={[
                    {name: 'qwr', url: 'http://localhost:8080'},
                    {name: 'qwr', url: 'http://localhost:8080'},
                    {name: 'qwr', url: 'http://localhost:8080'},
                ]}/>
            </main>
        </>
    )
}

export default App
