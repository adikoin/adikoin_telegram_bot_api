(function($) {

  "use strict";

  // Tab Section
  var initTabs = function(){
    const tabs = document.querySelectorAll('[data-tab-target]')
    const tabContents = document.querySelectorAll('[data-tab-content]')

    tabs.forEach(tab => {
      tab.addEventListener('click', () => {
        const target = document.querySelector(tab.dataset.tabTarget)
        tabContents.forEach(tabContent => {
          tabContent.classList.remove('active')
        })
        tabs.forEach(tab => {
          tab.classList.remove('active')
        })
        tab.classList.add('active')
        target.classList.add('active')
      })
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


// document ready
  $(document).ready(function(){
  
    let tg = window.Telegram.WebApp;
    tg.expand();

    initTabs();

    window.addEventListener("load", (event) => {
      //isotope
      $('.isotope-container').isotope({
        // options
        itemSelector: '.item',
        layoutMode: 'masonry'
      });



      // Initialize Isotope
      var $container = $('.isotope-container').isotope({
        // options
        itemSelector: '.item',
        layoutMode: 'masonry'
      });

      $(document).ready(function () {
        //active button
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

    // console.log("user:", tg.initDataUnsafe.user)


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

          // let formData = $(this).serialize(); // Собрать данные формы
          // console.log("isBot:", tg)
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
  });

})(jQuery);