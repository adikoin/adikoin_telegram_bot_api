(function($) {

  "use strict";

  // Tab Section
  var initTabs = function(){
    const tabs = document.querySelectorAll('[data-tab-target]');
    const tabContents = document.querySelectorAll('[data-tab-content]');

    tabs.forEach(tab => {
      tab.addEventListener('click', () => {
        const target = document.querySelector(tab.dataset.tabTarget);
        tabContents.forEach(tabContent => {
          tabContent.classList.remove('active');
        });
        tabs.forEach(tab => {
          tab.classList.remove('active');
        });
        tab.classList.add('active');
        target.classList.add('active');
      });
    });
  }

  var swiper = new Swiper(".product-swiper", {
    slidesPerView: 3,
    spaceBetween: 50,
    loop: true,
    pagination: {
      el: ".swiper-pagination",
      clickable: true,
    },
    breakpoints: {
      0: {
        slidesPerView: 1,
        spaceBetween: 20,
      },
      699: {
        slidesPerView: 2,
        spaceBetween: 30,
      },
      1200: {
        slidesPerView: 3,
        spaceBetween: 50,
      },
    },
  });

  // Show spinner
  function showSpinner() {
    document.getElementById('spinner').classList.remove('hidden');
    document.getElementById('user-info').classList.add('hidden');
    document.getElementById('early-adopters-form').classList.add('hidden');
    document.getElementById('error-container').classList.add('hidden');
  }

  // Hide spinner
  function hideSpinner() {
    document.getElementById('spinner').classList.add('hidden');
  }

  // Show user info
  function showUserInfo(data) {
    let tg = window.Telegram.WebApp;
    let telegramUser = tg.initDataUnsafe.user;
    const userInfoContainer = document.getElementById('user-info');
    const userNameElement = document.getElementById('user-name');
    const userBalanceElement = document.getElementById('user-balance');
    const userDurationElement = document.getElementById('user-duration');
    
    // Set the user's Telegram username
    userNameElement.innerText = telegramUser.username;
    
    // Extract and format the balance
    const amount = data.balance;
    userBalanceElement.innerText = formatNumber(amount);
    
    const remainingTimeStr = data.remaining;
  const remainingTimeMs = parseRemainingTime(remainingTimeStr);

  if (remainingTimeMs < 0) {
    // If the remaining time is negative, show the claim button
    userDurationElement.innerHTML = '<button id="claim-button">Claim</button>';
  } else {
    // Otherwise, display the remaining time in a user-friendly format
    userDurationElement.innerText = formatDuration(remainingTimeMs);
  }
   
    
    // Show the user info container
    userInfoContainer.classList.remove('hidden');
  }

  function parseRemainingTime(timeStr) {
    const isNegative = timeStr.startsWith('-');
    const timeParts = timeStr.match(/-?(\d+)m(\d+\.\d+)s/);
    if (timeParts) {
      const minutes = parseInt(timeParts[1], 10);
      const seconds = parseFloat(timeParts[2]);
      const totalMilliseconds = (minutes * 60 + seconds) * 1000;
      return isNegative ? -totalMilliseconds : totalMilliseconds;
    }
    return 0;
  }
  
  function formatDuration(durationMs) {
    const totalSeconds = Math.floor(durationMs / 1000);
    const minutes = Math.floor(totalSeconds / 60);
    const seconds = totalSeconds % 60;
    return `${minutes}m ${seconds}s`;
  }


  // Show form
  function showForm() {
    document.getElementById('early-adopters-form').classList.remove('hidden');
  }

  // Show error message
  function showError() {
    document.getElementById('error-container').classList.remove('hidden');
  }

  function formatNumber(num) {
    // Convert the number to string and split into integer and fractional parts
    let [integerPart, fractionalPart] = num.toString().split('.');
    // Separate the last three digits with a dot
    let lastThreeDigits = integerPart.slice(-3);
    let restOfTheNumber = integerPart.slice(0, -3);
    // Format the rest of the number with commas
    let formattedRest = restOfTheNumber.replace(/\B(?=(\d{3})+(?!\d))/g, ',');
    // Combine the formatted parts
    let formattedNumber = formattedRest ? formattedRest + '.' + lastThreeDigits : lastThreeDigits;
    return fractionalPart ? formattedNumber + '.' + fractionalPart : formattedNumber;
}

  // Fetch data from server
  async function fetchData(userId) {
    showSpinner();
    try {
      const response = await fetch(`https://tidy-mutually-reptile.ngrok-free.app/api/v1/early/${userId}`);
      
      if (!response.ok) {
        if (response.status === 302) {
          const data = await response.json();
          showUserInfo(data);
        } else if (response.status === 404) {
          showForm();
        } else {
          throw new Error(`Unexpected error: ${response.status}`);
        }
      } else {
        const data = await response.json();
        showUserInfo(data);
      }
    } catch (error) {
      console.error('Error fetching data:', error);
      showError();
    } finally {
      hideSpinner();
    }
  }

  // Document ready
  $(document).ready(function() {
    
    let tg = window.Telegram.WebApp;
    tg.expand();

    initTabs();

    window.addEventListener("load", (event) => {
      // Isotope initialization
      $('.isotope-container').isotope({
        itemSelector: '.item',
        layoutMode: 'masonry'
      });

      // Initialize Isotope
      var $container = $('.isotope-container').isotope({
        itemSelector: '.item',
        layoutMode: 'masonry'
      });

      $(document).ready(function () {
        // Active button
        $('.filter-button').click(function () {
          $('.filter-button').removeClass('active');
          $(this).addClass('active');
        });
      });

      // Filter items on button click
      $('.filter-button').click(function () {
        var filterValue = $(this).attr('data-filter');
        if (filterValue === '*') {
          // Show all items
          $container.isotope({ filter: '*' });
        } else {
          // Show filtered items
          $container.isotope({ filter: filterValue });
        }
      });
    });

    let telegramUser = tg.initDataUnsafe.user;
    fetchData(telegramUser.id);  // Fetch data when the document is ready, passing the user ID

    $('#form1').on('submit', function(event) {
      event.preventDefault(); // Остановить стандартную отправку формы

      // Скрыть все сообщения об ошибках
      $('.error-message').hide();

      // Валидация email
      let emailInput = $('#email1');
      let email = emailInput.val();
      let emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      
      if (!emailPattern.test(email)) {
        emailInput.addClass('has-error');
        $('#error-invalid-email').show();
      } else {
        emailInput.removeClass('has-error');

        let formData = {
          email: email,
          telegramUserID: telegramUser.id,
          isBot: telegramUser.is_bot,
          firstName: telegramUser.first_name,
          lastName: telegramUser.last_name || '',
          username: telegramUser.username,
          languageCode: telegramUser.language_code,
          isPremium: telegramUser.is_premium,
          addedToAttachmentMenu: telegramUser.added_to_attachment_menu,
          allowsWriteToPm: telegramUser.allows_write_to_pm,
          photoUrl: telegramUser.photo_url,
        };

        $.ajax({
          url: $(this).attr('action'), // URL для отправки
          type: $(this).attr('method'), // Метод отправки
          data: JSON.stringify(formData),
          contentType: 'application/json', 
          success: function(response, status, xhr) {
            if (xhr.status === 201) {
              // Успешное сохранение, код 201
              $('#form1').replaceWith('<div class="success-message">Спасибо за подписку!</div>');
            } else {
              // Обработка других успешных ответов (необязательно)
            }
          },
          error: function(xhr, status, error) {
            // Обработка ошибок
            if (xhr.status === 409) {
              $('#error-email-registered').show(); // Email уже зарегистрирован
            } else {
              $('#error-general').show(); // Что-то пошло не так
            }
          }
        });
      }
    });

    $('#start-farming-btn').on('click', function() {
      // Получение id из объекта window.Telegram.WebApp
      let id = tg.initDataUnsafe.user.id;

      // Формирование URL для POST запроса
      let url = `/api/v1/start_farm/${id}`;

      // Настройки для POST запроса
      let options = {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json'
          },
          body: JSON.stringify({ user_id: id })
      };

      // Отправка POST запроса с использованием fetch
      fetch(url, options)
          .then(response => response.json())
          .then(data => {
              console.log('Success:', data);
          })
          .catch((error) => {
              console.error('Error:', error);
          });
  });
  });

})(jQuery);
