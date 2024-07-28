// import { ReactComponent as MyIcon } from './path/to/my-icon.svg';

import TagEmoji from './tags';

function Card({
  title,
  reviews: { food, place },
  location: { googleMaps },
  tags,
}) {
  return (
    <div className="block max-w-sm p-6 bg-white border border-gray-200 rounded-lg dark:bg-gray-800 dark:border-gray-700 h-screen">
      <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
        {title}
      </h5>
      <div class="flex items-center justify-center mt-5 mb-5">
        <a href={googleMaps} class="text-sm text-gray-600 flex items-center">
          <Direction />
        </a>
      </div>
      <div className="flex items-center justify-center space-x-4  mt-5 mb-5">
        {tags.map((tag, i) => (
          <TagEmoji tag={tag} key={i} />
        ))}
      </div>

      {food !== '' && place !== '' && (
        <>
          <h4 class="mb-2 text-lg font-bold tracking-tight text-gray-900 dark:text-white">
            Food & Place
          </h4>
          <p class="font-normal text-gray-700 dark:text-gray-400">
            <div dangerouslySetInnerHTML={{ __html: food }} />
          </p>
          <br />
          <p class="font-normal text-gray-700 dark:text-gray-400">
            <div dangerouslySetInnerHTML={{ __html: place }} />
          </p>
        </>
      )}
    </div>
  );
}

function Direction() {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="#fff"
      stroke="currentColor"
      stroke-width="2"
      stroke-linecap="round"
      stroke-linejoin="round"
      class="feather feather-compass"
    >
      <circle cx="12" cy="12" r="10"></circle>
      <polygon points="16.24 7.76 14.12 14.12 7.76 16.24 9.88 9.88 16.24 7.76"></polygon>
    </svg>
  );
}

export default Card;
