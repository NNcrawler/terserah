import './App.css';
import LocationFetcher from './LocationFetcher';

function App() {
  return (
    <div className="App">
      <div class="flex items-center justify-center h-screen">
        <div>
          <LocationFetcher />
        </div>
      </div>
    </div>
  );
}

export default App;
