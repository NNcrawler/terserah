import React from 'react';
import Slider from 'react-slick';
import 'slick-carousel/slick/slick.css';
import 'slick-carousel/slick/slick-theme.css';

export default function SimpleSlider({ children }) {
  var settings = {
    infinite: true,
    slidesToShow: 1,
    slidesToScroll: 1,
    vertical: true,
    verticalSwiping: true,
    adaptiveHeight: true,
  };
  return (
    <Slider className="flex items-center justify-center" {...settings}>
      {children}
    </Slider>
  );
}