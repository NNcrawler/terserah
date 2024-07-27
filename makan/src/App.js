import { useEffect, useState } from 'react';
import fetchRecommendations from './api/recommendation';
import './App.css';
import LocationFetcher from './LocationFetcher';
import { useGeolocation } from '@uidotdev/usehooks';

function App() {
  const location = useGeolocation();

  const [recommendations, setRecommendations] = useState([]);
  useEffect(() => {
    if (location.latitude && location.longitude) {
      fetchRecommendations(location).then((response) => {
        setRecommendations(response.data);
      });
    }
  }, [location]);

  return (
    <div className="App">
      <div class="flex items-center justify-center h-screen">
        <div>
          <LocationFetcher location={location} />
          {recommendations.length > 0 &&
            recommendations.map((recommendation, i) => (
              <p id={i}>{recommendation.name}</p>
            ))}
        </div>
      </div>
    </div>
  );
}

export default App;
