import './App.css';

import { useGeolocation } from '@uidotdev/usehooks';
import { useEffect, useState } from 'react';
import fetchRecommendations from './api/recommendation';
import Card from './components/card';
import SlideShow from './components/slider';

function App() {
  const location = useGeolocation();

  const [recommendations, setRecommendations] = useState([]);
  useEffect(() => {
    const latitude = location.latitude;
    const longitude = location.longitude;
    if (latitude && longitude) {
      fetchRecommendations({
        longitude: longitude,
        latitude: latitude,
      }).then((response) => {
        setRecommendations(response.data);
      });
    }
  }, [location]);


  return (
    <div className="App">
      <div className="flex items-center justify-center h-screen">
        <div className="max-w-sm h-screen">
          {recommendations.length > 0 && (
            <SlideShow slides={recommendations}>
              {recommendations.map((recommendation, i) => (
                <Card
                  key={i}
                  tags = {recommendation.tags}
                  title={recommendation.name}
                  reviews={recommendation.reviews}
                  location={recommendation.location}
                  dishes={recommendation.dishType}
                  distance={recommendation.distance}
                />
              ))}
            </SlideShow>
          )}
        </div>
      </div>
    </div>
  );
}

export default App;
