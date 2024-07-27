import './App.css';
import LocationFetcher from './LocationFetcher';
import { useGeolocation } from '@uidotdev/usehooks';

function App() {
  const location = useGeolocation();

  return (
    <div className="App">
      <div class="flex items-center justify-center h-screen">
        <div>
          <LocationFetcher location={location} />
        </div>
      </div>
    </div>
  );
}

export default App;
