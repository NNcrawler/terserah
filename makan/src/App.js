import './App.css';

import { useGeolocation } from '@uidotdev/usehooks';
import copyWriter from './api/copywriter';
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

  const [reviews, setReviews] = useState([]);
  useEffect(() => {
    if (recommendations.length > 0) {
      copyWriter('reviewSummarizer', {
        reviews: recommendations[0].reviews,
      }).then((reviews) => {
        setReviews(reviews);
      });
    }
    console.log(recommendations);
  }, [recommendations]);

  return (
    <div className="App">
      <div className="flex items-center justify-center h-screen">
        <div className="max-w-sm h-screen">
          {reviews && (
            <SlideShow slides={recommendations}>
              {recommendations.map((recommendation, i) => (
                <Card
                  key={i}
                  title={recommendation.name}
                  reviews={recommendation.reviews}
                  location={recommendation.location}
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
